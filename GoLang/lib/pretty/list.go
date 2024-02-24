package pretty // import "my-module/lib/pretty"

import "fmt"

// listInt print list of ints separated by commas
func listInt (a []int) string {
  var items string
  items += "["
  for _, element := range a {
    if l := len(a); element == a[l - 1] {
      items += fmt.Sprintf(`'%v'`, element)
      continue
    }
    items += fmt.Sprintf(`'%v', `, element)
  }
  items += "]"
  return items
}

// List create a pretty list and return it as a string
func List (i interface{}) string {
  switch i := i.(type) {
    case []int:
      return listInt(i)
    case []string:
      str := "["
      for _, e := range i {
        if l := len(i); e == i[l - 1] {
          str += fmt.Sprintf(`'%v'`, e)
          continue
        }
        str += fmt.Sprintf(`'%v', `, e)
      }
      str += "]"
      return str
    default:
      return fmt.Sprintf("unsupported type: %T\n", i)
  }
}
