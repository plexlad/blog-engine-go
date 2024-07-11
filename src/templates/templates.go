package templates

import (
	"html/template"
	"io"
	"os"
)

// Renders views using a template
// templates is a hashmap of templates used for rendering
//
type ViewRenderer struct {
  templates map[string]*template.Template
  base_name string
}

// TODO: Error handling

// Use data to render templates
func (vr *ViewRenderer) RenderWithData(w io.Writer, file_name string, data interface{}) error {
  return vr.templates[file_name].ExecuteTemplate(w, vr.base_name, data)
}

// Templates to render that don't require data
func (vr *ViewRenderer) RenderWithoutData(w io.Writer, file_name string) error {
  return vr.templates[file_name].ExecuteTemplate(w, vr.base_name, nil)
}

// Creates a ViewRenderer from a file string
// directory is the directory of the template files
// base_file is the file base 
// base_name is the name of the base in the template (MAKE SURE THIS IS CORRECT)
func ViewRendererFromFilePath(directory_path string, base_file_path string, base_name string) (*ViewRenderer, error) {
  var new_renderer ViewRenderer

  // Set a directory to an os.File
  directory_file, err := os.Open(directory_path)
  if err != nil {
    return nil, err
  }

  // Loads the list of files in the directory
  files, err := directory_file.Readdir(0)
  if err != nil {
    return nil, err
  }
  
  // Loops through list of files; if file is a directory, skip
  for _, file_info := range files {
    if file_info.IsDir() {
      continue
    }

    new_renderer.templates[file_info.Name()] = *template.Must(template.ParseFiles(base_file_path, directory_path + file_info.Name()))
  }

  new_renderer.base_name = base_name

  return &new_renderer, nil
}
