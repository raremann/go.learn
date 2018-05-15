// Surface computes an SVG rendering of a 3-D surface function
package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320            // canvas size, pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges
	zscale        = height * 0.4        // pixels per z unit
	xyscale       = width / 2 / xyrange // pixels per unit
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

/*
func main() {
	surface(os.Stdout)
}
*/
func Surface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aok := corner(i+1, j)
			bx, by, bok := corner(i, j)
			cx, cy, cok := corner(i, j+1)
			dx, dy, dok := corner(i+1, j+1)
			col := getColor(i, j)
			if aok && bok && cok && dok {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n", ax, ay,
					bx, by, cx, cy, dx, dy, col)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

// Compute a color for a cell based on the height
func getColor(i, j int) string {
	// Hack to hardcode max and min z values
	// zmax := 1.0
	// zmin := -1.0
	_, _, z := coordinates(i, j)
	result := "#0000ff"
	if z > 0 {
		result = "#ff0000"
	}
	return result
}

// Find the cartesian coordinates of the corner and it's height
func coordinates(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute the surface height.
	z := f(x, y)
	return x, y, z
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell(i,j).
	x, y, z := coordinates(i, j)
	//x := xyrange * (float64(i)/cells - 0.5)
	//y := xyrange * (float64(j)/cells - 0.5)

	// Compute the surface height.
	//z := f(x, y)

	// Project (x,y,z) onto 2-d SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	ok := true
	if math.IsNaN(sx) || math.IsNaN(sy) {
		ok = false
	}
	return sx, sy, ok
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from origin
	return math.Sin(r) / r
}
