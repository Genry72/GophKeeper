package tuiclient

import (
	"context"
	"errors"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

func (a *App) tree(ctx context.Context) {
	rootDir := "./"
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		f, err := os.Open(path)
		if err != nil {
			if errors.Is(err, os.ErrPermission) {
				return
			}
		}

		fi, err := f.Stat()
		if err != nil {
			return
		}

		if !fi.IsDir() {
			a.formAddBinary(ctx, path, nil)
			a.tvievApp.pages.SwitchToPage(pageAny)

			return
		}

		files, err := os.ReadDir(path)
		if err != nil {
			return
		}

		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name()))

			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}

			target.AddChild(node)
		}
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	*a.tvievApp.treeView = *tree
}
