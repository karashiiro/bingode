package bingode

import "github.com/xivapi/godestone/v2"

type bingodeProvider struct{}

// New constructs a binary-data-backed data provider for use in
// a godestone parser.
func New() godestone.DataProvider {
	b := bingodeProvider{}
	return &b
}

func (b *bingodeProvider) Achievement(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) ClassJob(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Deity(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) GrandCompany(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Item(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Minion(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Mount(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Race(name string) *godestone.GenderedEntity {
	return nil
}

func (b *bingodeProvider) Reputation(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Title(name string) *godestone.Title {
	return nil
}

func (b *bingodeProvider) Town(name string) *godestone.NamedEntity {
	return nil
}

func (b *bingodeProvider) Tribe(name string) *godestone.GenderedEntity {
	return nil
}
