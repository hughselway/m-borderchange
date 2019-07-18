package grpc

import(
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	pbborderchange "videoseries/m-apis/m-borderchange"
)

type Grpc struct {
	address string
	srv *grpc.Server
}


func (g *Grpc) GetBorder(ctx context.Context, request pbborderchange.GetBorderRequest) (*pbborderchange.GetBorderResponse,error) {
	log.Info().Msg("GetBorder is called")
	prevstyle := request.Style
	prevwidth := request.Width
	shapeR :=  request.ShapeColorRed
	shapeG :=  request.ShapeColorGreen
	shapeB :=  request.ShapeColorBlue

	widths := [3]string{"thin","medium","thick"}
	styles := []string{"dotted","dashed","solid","double","groove","ridge"}

}