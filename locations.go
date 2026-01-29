package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)


type LocationAreaResponce struct {
	Results []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type LocationAreaPaginator struct {
	Cache map[int]LocationArea
	Size int
	Current int
	Limit int
}

func (LAP *LocationAreaPaginator) Init(limit int) {
	LAP.Limit = limit
	LAP.Current = 0
	LAP.Size = 0
	LAP.Cache = make(map[int]LocationArea)
}

func (LAP *LocationAreaPaginator) GrabPage(offset int)([]LocationArea, error) {
	laList := make([]LocationArea, 0, 20)
	for i := 0; i < LAP.Limit; i ++ {
		v, ok := LAP.Cache[offset + i];
		if !ok {
			err := LAP.getLocations(offset + i, LAP.Limit)
			if err != nil {
				return nil, err
			}
			v, ok = LAP.Cache[offset + i]
			if !ok {
				return nil, fmt.Errorf("Cannot retrieve data for %d location", offset + i)
			}
		}
		laList = append(laList, v)
	}
	return laList, nil
}

func (LAP *LocationAreaPaginator) NextPage()([]LocationArea, error) {
	laList, err := LAP.GrabPage(LAP.Current + LAP.Limit)
	if err != nil {
		return nil, err
	} else {
		LAP.Current = LAP.Current + LAP.Limit
		return laList, err
	}
}
func (LAP *LocationAreaPaginator) PrevPage()([]LocationArea, error) {
	laList, err := LAP.GrabPage(LAP.Current - LAP.Limit)
	if err != nil {
		return nil, err
	} else {
		LAP.Current = LAP.Current - LAP.Limit
		if LAP.Current < 0 {
			LAP.Current = 0
		}
		return laList, err
	}
}

func (LAP *LocationAreaPaginator) getLocations(offset, limit int) error {
	if offset < 0 {
		offset = 0
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=%d", offset, limit)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch location areas")
	}
	var data LocationAreaResponce
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	} else {
		for i, location := range data.Results {
			_, ok := LAP.Cache[offset + i]
			if !ok {
				LAP.Size ++
			}
			LAP.Cache[offset + i] = location
		}
	}
	return nil
}

func (LAP *LocationAreaPaginator) pruneCache() {
	// prune cache down to 3 * limit size
}
