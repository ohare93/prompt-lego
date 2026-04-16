package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// titleCase converts "some-name" to "Some-Name" (ASCII only).
func titleCase(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		if len(p) == 0 {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "-")
}

var version = "dev"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	showVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return nil
	}

	dirs := collectDirs()
	s := server.NewMCPServer("prompts", version)

	seen := map[string]bool{}
	for _, dir := range dirs {
		info, err := os.Stat(dir)
		if err != nil || !info.IsDir() {
			continue
		}
		project := findProjectName(dir)
		promptsDir := findPromptsDir(dir)
		files := scanPrompts(promptsDir)

		for _, f := range files {
			uri := project + "://" + strings.TrimSuffix(filepath.Base(f), ".md")
			if seen[uri] {
				continue
			}
			seen[uri] = true

			title, desc := extractMetadata(f)
			resource := mcp.NewResource(
				uri,
				title,
				mcp.WithResourceDescription(desc),
				mcp.WithMIMEType("text/markdown"),
			)
			path := f
			s.AddResource(resource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
				content, err := os.ReadFile(path)
				if err != nil {
					return nil, err
				}
				return []mcp.ResourceContents{
					mcp.TextResourceContents{
						URI:      uri,
						MIMEType: "text/markdown",
						Text:     string(content),
					},
				}, nil
			})
		}
	}

	return server.ServeStdio(s)
}

func collectDirs() []string {
	var dirs []string
	seen := map[string]bool{}

	add := func(p string) {
		abs, err := filepath.Abs(p)
		if err != nil {
			return
		}
		if !seen[abs] {
			seen[abs] = true
			dirs = append(dirs, abs)
		}
	}

	// From PROMPT_DIRS env var
	if env := os.Getenv("PROMPT_DIRS"); env != "" {
		for _, d := range strings.Split(env, ",") {
			d = strings.TrimSpace(d)
			if d != "" {
				add(d)
			}
		}
	}

	// From command-line arguments
	for _, arg := range flag.Args() {
		add(arg)
	}

	// Auto-detect prompts/ in cwd
	cwd, err := os.Getwd()
	if err == nil {
		if info, err := os.Stat(filepath.Join(cwd, "prompts")); err == nil && info.IsDir() {
			abs, _ := filepath.Abs(cwd)
			if !seen[abs] {
				dirs = append([]string{abs}, dirs...)
				seen[abs] = true
			}
		}
	}

	return dirs
}

func findProjectName(dir string) string {
	current := dir
	for {
		gitPath := filepath.Join(current, ".git")
		if _, err := os.Stat(gitPath); err == nil {
			return filepath.Base(current)
		}
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}
	return filepath.Base(dir)
}

func findPromptsDir(dir string) string {
	p := filepath.Join(dir, "prompts")
	if info, err := os.Stat(p); err == nil && info.IsDir() {
		return p
	}
	return dir
}

func scanPrompts(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		files = append(files, filepath.Join(dir, e.Name()))
	}
	sort.Strings(files)
	return files
}

func extractMetadata(path string) (title, description string) {
	f, err := os.Open(path)
	if err != nil {
		name := filepath.Base(path)
		name = strings.TrimSuffix(name, ".md")
		return titleCase(name), ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	first := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if first {
			first = false
			if strings.HasPrefix(line, "# ") {
				title = strings.TrimPrefix(line, "# ")
				continue
			}
			name := filepath.Base(path)
			name = strings.TrimSuffix(name, ".md")
			title = titleCase(name)
		}
		if !strings.HasPrefix(line, "#") {
			if len(line) > 200 {
				description = line[:200]
			} else {
				description = line
			}
			return
		}
	}

	if title == "" {
		name := filepath.Base(path)
		name = strings.TrimSuffix(name, ".md")
		title = titleCase(name)
	}
	return
}
