package grpc

import (
	"context"
	pbborderchange "github.com/hughselway/m-apis/m-borderchange"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"math/rand"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetBorder(ctx context.Context, request *pbborderchange.GetBorderRequest) (*pbborderchange.GetBorderResponse, error) {
	log.Info().Msg("GetBorder is called")

	widths := [3]string{"thin", "medium", "thick"}
	styles := [8]string{"dotted", "dashed", "solid", "double", "groove", "ridge", "inset", "outset"}
	var outputstyle string
	var outputwidth string

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))

	outputwidth = widths[rand.Intn(3)]
	outputstyle = styles[rand.Intn(8)]

	//for i := range widths {
	//	if (*request).Width == widths[i] {
	//		outputwidth = widths[(i+1)%3]
	//		break
	//	}
	//	return &pbborderchange.GetBorderResponse{}, errors.New("requested width is not accepted")
	//}

	//for i := range styles {
	//	if (*request).Style == styles[i] {
	//		outputstyle = styles[(i+1)%8]
	//		break
	//	}
	//	return &pbborderchange.GetBorderResponse{}, errors.New("requested style is not accepted")
	//}

	return &pbborderchange.GetBorderResponse{
		Style:            outputstyle,
		Width:            outputwidth,
		BorderColorRed:   256 * rand.Float64(),
		BorderColorGreen: 256 * rand.Float64(),
		BorderColorBlue:  256 * rand.Float64(),
		//BorderColorRed:   255 - (*request).ShapeColorRed,
		//BorderColorGreen: 255 - (*request).ShapeColorGreen,
		//BorderColorBlue:  255 - (*request).ShapeColorBlue,
	}, nil

}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)

	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	g.srv = grpc.NewServer()

	pbborderchange.RegisterGameBorderServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting grpc server for borderchange microservice")

	err = g.srv.Serve(lis)

	if err != nil {
		return errors.Wrap(err, "failed to start grpc server for borderchange microservice")
	}
	return nil
}
