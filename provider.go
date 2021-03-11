package bingode

import "github.com/karashiiro/godestone"

type bingodeProvider struct{}

// New constructs a binary-data-backed data provider for use in
// a godestone parser.
func New() *godestone.DataProvider {
	b := &bingodeProvider{}
	return b
}

func (b *bingodeProvider) ClassJob(name string) models.ClassJob        {}
func (b *bingodeProvider) Deity(name string) models.NamedEntity        {}
func (b *bingodeProvider) GrandCompany(name string) models.NamedEntity {}
func (b *bingodeProvider) Item(name string) exports.Item               {}
func (b *bingodeProvider) Minion(name string) models.Minion            {}
func (b *bingodeProvider) Mount(name string) models.Mount              {}
func (b *bingodeProvider) Race(name string) models.GenderedEntity      {}
func (b *bingodeProvider) Reputation(name string) models.NamedEntity   {}
func (b *bingodeProvider) Title(name string) models.Title              {}
func (b *bingodeProvider) Town(name string) models.NamedEntity         {}
func (b *bingodeProvider) Tribe(name string) models.GenderedEntity     {}
