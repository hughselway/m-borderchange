package grpc

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	pbborderchange "videoseries/m-apis/m-borderchange"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func (g *Grpc) GetBorder(ctx context.Context, request *pbborderchange.GetBorderRequest) (*pbborderchange.GetBorderResponse, error) {
	log.Info().Msg("GetBorder is called")

	widths := [3]string{"thin", "medium", "thick"}
	styles := [8]string{"dotted", "dashed", "solid", "double", "groove", "ridge", "inset", "outset"}
	var outputstyle string
	var outputwidth string

	for i := range widths {
		if (*request).Width == widths[i] {
			outputwidth = widths[(i+1)%3]
			break
		}
		return &pbborderchange.GetBorderResponse{}, errors.New("requested width is not accepted")
	}

	for i := range styles {
		if (*request).Style == styles[i] {
			outputstyle = styles[(i+1)%8]
			break
		}
		return &pbborderchange.GetBorderResponse{}, errors.New("requested style is not accepted")
	}

	return &pbborderchange.GetBorderResponse{
		Style:            outputstyle,
		Width:            outputwidth,
		BorderColorRed:   255 - (*request).ShapeColorRed,
		BorderColorGreen: 255 - (*request).ShapeColorGreen,
		BorderColorBlue:  255 - (*request).ShapeColorBlue,
	}, nil

}
