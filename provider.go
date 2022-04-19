package bingode

import (
	"errors"
	"strings"

	"github.com/karashiiro/bingode/internal/pack/exports"

	"github.com/xivapi/godestone/v2/provider"
	"github.com/xivapi/godestone/v2/provider/models"
)

const notFound = "no entity matching the criteria was found"

type bingodeProvider struct {
	achievementTable  *exports.AchievementTable
	classJobTable     *exports.ClassJobTable
	deityTable        *exports.DeityTable
	grandCompanyTable *exports.GrandCompanyTable
	itemTable         *exports.ItemTable
	minionTable       *exports.MinionTable
	mountTable        *exports.MountTable
	raceTable         *exports.RaceTable
	repTable          *exports.ReputationTable
	titleTable        *exports.TitleTable
	townTable         *exports.TownTable
	tribeTable        *exports.TribeTable
}

// New constructs a binary-data-backed data provider for use in
// a godestone parser.
func New() provider.DataProvider {
	b := bingodeProvider{}
	return &b
}

func (b *bingodeProvider) getAchievementTable() *exports.AchievementTable {
	if b.achievementTable == nil {
		data, _ := exports.Asset("achievement_table.bin")
		achievementTable := exports.GetRootAsAchievementTable(data, 0)
		b.achievementTable = achievementTable
	}
	return b.achievementTable
}

func (b *bingodeProvider) getClassJobTable() *exports.ClassJobTable {
	if b.classJobTable == nil {
		data, _ := exports.Asset("classjob_table.bin")
		classJobTable := exports.GetRootAsClassJobTable(data, 0)
		b.classJobTable = classJobTable
	}
	return b.classJobTable
}

func (b *bingodeProvider) getDeityTable() *exports.DeityTable {
	if b.deityTable == nil {
		data, _ := exports.Asset("deity_table.bin")
		deityTable := exports.GetRootAsDeityTable(data, 0)
		b.deityTable = deityTable
	}
	return b.deityTable
}

func (b *bingodeProvider) getGrandCompanyTable() *exports.GrandCompanyTable {
	if b.grandCompanyTable == nil {
		data, _ := exports.Asset("gc_table.bin")
		grandCompanyTable := exports.GetRootAsGrandCompanyTable(data, 0)
		b.grandCompanyTable = grandCompanyTable
	}
	return b.grandCompanyTable
}

func (b *bingodeProvider) getItemTable() *exports.ItemTable {
	if b.itemTable == nil {
		data, _ := exports.Asset("item_table.bin")
		itemTable := exports.GetRootAsItemTable(data, 0)
		b.itemTable = itemTable
	}
	return b.itemTable
}

func (b *bingodeProvider) getMinionTable() *exports.MinionTable {
	if b.minionTable == nil {
		data, _ := exports.Asset("minion_table.bin")
		minionTable := exports.GetRootAsMinionTable(data, 0)
		b.minionTable = minionTable
	}
	return b.minionTable
}

func (b *bingodeProvider) getMountTable() *exports.MountTable {
	if b.mountTable == nil {
		data, _ := exports.Asset("mount_table.bin")
		mountTable := exports.GetRootAsMountTable(data, 0)
		b.mountTable = mountTable
	}
	return b.mountTable
}

func (b *bingodeProvider) getRaceTable() *exports.RaceTable {
	if b.raceTable == nil {
		data, _ := exports.Asset("race_table.bin")
		raceTable := exports.GetRootAsRaceTable(data, 0)
		b.raceTable = raceTable
	}
	return b.raceTable
}

func (b *bingodeProvider) getReputationTable() *exports.ReputationTable {
	if b.repTable == nil {
		data, _ := exports.Asset("reputation_table.bin")
		repTable := exports.GetRootAsReputationTable(data, 0)
		b.repTable = repTable
	}
	return b.repTable
}

func (b *bingodeProvider) getTitleTable() *exports.TitleTable {
	if b.titleTable == nil {
		data, _ := exports.Asset("title_table.bin")
		titleTable := exports.GetRootAsTitleTable(data, 0)
		b.titleTable = titleTable
	}
	return b.titleTable
}

func (b *bingodeProvider) getTownTable() *exports.TownTable {
	if b.townTable == nil {
		data, _ := exports.Asset("town_table.bin")
		townTable := exports.GetRootAsTownTable(data, 0)
		b.townTable = townTable
	}
	return b.townTable
}

func (b *bingodeProvider) getTribeTable() *exports.TribeTable {
	if b.tribeTable == nil {
		data, _ := exports.Asset("tribe_table.bin")
		tribeTable := exports.GetRootAsTribeTable(data, 0)
		b.tribeTable = tribeTable
	}
	return b.tribeTable
}

func (b *bingodeProvider) Achievement(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getAchievementTable().AchievementsLength()
	for i := 0; i < nLength; i++ {
		o := exports.Achievement{}
		b.getAchievementTable().Achievements(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) ClassJob(name string) (*models.ClassJobInternal, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getClassJobTable().ClassJobsLength()
	for i := 0; i < nLength; i++ {
		o := exports.ClassJob{}
		b.getClassJobTable().ClassJobs(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.ClassJobInternal{
				NamedEntity: &models.NamedEntity{
					ID:   o.Id(),
					Name: name,

					NameEN: nameEn,
					NameDE: nameDe,
					NameFR: nameFr,
					NameJA: nameJa,
				},

				Parent:   o.Parent(),
				JobIndex: o.JobIndex(),
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) JobForClass(name string) (*models.ClassJobInternal, error) {
	cj, err := b.ClassJob(name)
	if err != nil {
		return nil, err
	}

	nLength := b.getClassJobTable().ClassJobsLength()
	for i := 0; i < nLength; i++ {
		o := exports.ClassJob{}
		b.getClassJobTable().ClassJobs(&o, i)

		if o.JobIndex() != 0 {
			continue
		}

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())
		if cj.Parent == o.Parent() {
			return &models.ClassJobInternal{
				NamedEntity: &models.NamedEntity{
					ID:   o.Id(),
					Name: name,

					NameEN: nameEn,
					NameDE: nameDe,
					NameFR: nameFr,
					NameJA: nameJa,
				},

				Parent:   o.Parent(),
				JobIndex: o.JobIndex(),
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Deity(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getDeityTable().DeitiesLength()
	for i := 0; i < nLength; i++ {
		o := exports.Deity{}
		b.getDeityTable().Deities(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) GrandCompany(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getGrandCompanyTable().GrandCompaniesLength()
	for i := 0; i < nLength; i++ {
		o := exports.GrandCompany{}
		b.getGrandCompanyTable().GrandCompanies(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Item(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getItemTable().ItemsLength()
	for i := 0; i < nLength; i++ {
		o := exports.Item{}
		b.getItemTable().Items(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Minion(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getMinionTable().MinionsLength()
	for i := 0; i < nLength; i++ {
		o := exports.Minion{}
		b.getMinionTable().Minions(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Mount(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getMountTable().MountsLength()
	for i := 0; i < nLength; i++ {
		o := exports.Mount{}
		b.getMountTable().Mounts(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Race(name string) (*models.GenderedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getRaceTable().RacesLength()
	for i := 0; i < nLength; i++ {
		o := exports.Race{}
		b.getRaceTable().Races(&o, i)

		nameMasculineEn := string(o.NameMasculineEn())
		nameMasculineJa := string(o.NameMasculineJa())
		nameMasculineDe := string(o.NameMasculineDe())
		nameMasculineFr := string(o.NameMasculineFr())
		nameFeminineEn := string(o.NameFeminineEn())
		nameFeminineJa := string(o.NameFeminineJa())
		nameFeminineDe := string(o.NameFeminineDe())
		nameFeminineFr := string(o.NameFeminineFr())

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineJa,
			nameMasculineDe,
			nameMasculineFr,
			nameFeminineEn,
			nameFeminineJa,
			nameFeminineDe,
			nameFeminineFr,
		) {
			return &models.GenderedEntity{
				ID:   o.Id(),
				Name: name,

				NameMasculineEN: nameMasculineEn,
				NameMasculineJA: nameMasculineJa,
				NameMasculineDE: nameMasculineDe,
				NameMasculineFR: nameMasculineFr,
				NameFeminineEN:  nameFeminineEn,
				NameFeminineJA:  nameFeminineJa,
				NameFeminineDE:  nameFeminineDe,
				NameFeminineFR:  nameFeminineFr,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Reputation(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getReputationTable().ReputationsLength()
	for i := 0; i < nLength; i++ {
		o := exports.Reputation{}
		b.getReputationTable().Reputations(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Title(name string) (*models.TitleInternal, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getTitleTable().TitlesLength()
	for i := 0; i < nLength; i++ {
		o := exports.Title{}
		b.getTitleTable().Titles(&o, i)

		nameMasculineEn := string(o.NameMasculineEn())
		nameMasculineJa := string(o.NameMasculineJa())
		nameMasculineDe := string(o.NameMasculineDe())
		nameMasculineFr := string(o.NameMasculineFr())
		nameFeminineEn := string(o.NameFeminineEn())
		nameFeminineJa := string(o.NameFeminineJa())
		nameFeminineDe := string(o.NameFeminineDe())
		nameFeminineFr := string(o.NameFeminineFr())

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineJa,
			nameMasculineDe,
			nameMasculineFr,
			nameFeminineEn,
			nameFeminineJa,
			nameFeminineDe,
			nameFeminineFr,
		) {
			ge := &models.GenderedEntity{
				ID:   o.Id(),
				Name: name,

				NameMasculineEN: nameMasculineEn,
				NameMasculineJA: nameMasculineJa,
				NameMasculineDE: nameMasculineDe,
				NameMasculineFR: nameMasculineFr,
				NameFeminineEN:  nameFeminineEn,
				NameFeminineJA:  nameFeminineJa,
				NameFeminineDE:  nameFeminineDe,
				NameFeminineFR:  nameFeminineFr,
			}

			return &models.TitleInternal{
				GenderedEntity: ge,
				Prefix:         o.IsPrefix(),
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Town(name string) (*models.NamedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getTownTable().TownsLength()
	for i := 0; i < nLength; i++ {
		o := exports.Town{}
		b.getTownTable().Towns(&o, i)

		nameEn := string(o.NameEn())
		nameDe := string(o.NameDe())
		nameFr := string(o.NameFr())
		nameJa := string(o.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &models.NamedEntity{
				ID:   o.Id(),
				Name: name,

				NameEN: nameEn,
				NameDE: nameDe,
				NameFR: nameFr,
				NameJA: nameJa,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}

func (b *bingodeProvider) Tribe(name string) (*models.GenderedEntity, error) {
	nameLower := strings.ToLower(name)

	nLength := b.getTribeTable().TribesLength()
	for i := 0; i < nLength; i++ {
		o := exports.Tribe{}
		b.getTribeTable().Tribes(&o, i)

		nameMasculineEn := string(o.NameMasculineEn())
		nameMasculineJa := string(o.NameMasculineJa())
		nameMasculineDe := string(o.NameMasculineDe())
		nameMasculineFr := string(o.NameMasculineFr())
		nameFeminineEn := string(o.NameFeminineEn())
		nameFeminineJa := string(o.NameFeminineJa())
		nameFeminineDe := string(o.NameFeminineDe())
		nameFeminineFr := string(o.NameFeminineFr())

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineJa,
			nameMasculineDe,
			nameMasculineFr,
			nameFeminineEn,
			nameFeminineJa,
			nameFeminineDe,
			nameFeminineFr,
		) {
			return &models.GenderedEntity{
				ID:   o.Id(),
				Name: name,

				NameMasculineEN: nameMasculineEn,
				NameMasculineJA: nameMasculineJa,
				NameMasculineDE: nameMasculineDe,
				NameMasculineFR: nameMasculineFr,
				NameFeminineEN:  nameFeminineEn,
				NameFeminineJA:  nameFeminineJa,
				NameFeminineDE:  nameFeminineDe,
				NameFeminineFR:  nameFeminineFr,
			}, nil
		}
	}

	return nil, errors.New(notFound)
}
