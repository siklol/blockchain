package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/siklol/blockchain"
)

func Version(p *Peer) (int, error) {
	v := struct {
		Version json.Number `json:"version"`
	}{}

	rsp, err := get(p, "/version")
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(rsp, &v); err != nil {
		return 0, err
	}

	i, err := v.Version.Int64()

	if err != nil {
		return 0, err
	}

	return int(i), nil
}

func GenesisBlock(p *Peer) (*blockchain.Block, error) {
	rsp, err := get(p, "/blocks/genesis")
	if err != nil {
		return nil, err
	}

	return block(rsp)
}

func Tip(p *Peer) (*blockchain.Block, error) {
	rsp, err := get(p, "/blocks/tip")
	if err != nil {
		return nil, err
	}

	return block(rsp)
}

func BlockAtIndex(p *Peer, i int64) (*blockchain.Block, error) {
	rsp, err := get(p, fmt.Sprintf("/blocks/index/%d", i))
	if err != nil {
		return nil, err
	}

	return block(rsp)
}

func block(rsp []byte) (*blockchain.Block, error) {
	var block *blockchain.Block
	if err := json.Unmarshal(rsp, &block); err != nil {
		return nil, err
	}

	// FIXME this is hardcoded. Change this

	return block, nil
}

func get(p *Peer, url string) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("http://%s:%s"+url, p.IP, p.Port)) // TODO https?
	if err != nil {
		return []byte(""), err
	}

	rsp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte(""), err
	}

	return rsp, nil
}