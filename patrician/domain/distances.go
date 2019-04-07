package domain

// Distances The distances between cities
var Distances = map[string]map[string]float64{
	"Edimburgo": map[string]float64{
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Scarborough": map[string]float64{
		"Edimburgo": 2,
		"Londres":   2,
		"Brujas":    2,
		"Colonia":   2,
		"Groninga":  2,
		"Bremen":    2,
		"Hamburgo":  2,
		"Ripen":     2,
		"Bergen":    2,
		"Oslo":      2,
		"Aalborg":   2,
		"Malmo":     2,
		"Lubeck":    2,
		"Rostock":   2,
		"Stettin":   2,
		"Gdansk":    2,
		"Torum":     2,
		"Riga":      2,
		"Visby":     2,
		"Estocolmo": 2,
		"Reval":     2,
		"Ladoga":    2,
		"Novgorod":  2},
	"Londres": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Brujas": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Colonia": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Groninga": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Bremen": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Hamburgo": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Ripen": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Bergen": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Oslo": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Aalborg": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Malmo": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Lubeck": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Rostock": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Stettin": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Gdansk": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Torum": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Riga": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Visby":       57,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Visby": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        57,
		"Estocolmo":   36,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Estocolmo": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       36,
		"Reval":       2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Reval": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Ladoga":      2,
		"Novgorod":    2},
	"Ladoga": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Novgorod":    2},
	"Novgorod": map[string]float64{
		"Edimburgo":   2,
		"Scarborough": 2,
		"Londres":     2,
		"Brujas":      2,
		"Colonia":     2,
		"Groninga":    2,
		"Bremen":      2,
		"Hamburgo":    2,
		"Ripen":       2,
		"Bergen":      2,
		"Oslo":        2,
		"Aalborg":     2,
		"Malmo":       2,
		"Lubeck":      2,
		"Rostock":     2,
		"Stettin":     2,
		"Gdansk":      2,
		"Torum":       2,
		"Riga":        2,
		"Visby":       2,
		"Estocolmo":   2,
		"Reval":       2,
		"Ladoga":      2},
}
