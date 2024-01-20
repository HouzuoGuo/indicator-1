// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator/v2

package strategy_test

import (
	"os"
	"testing"

	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/strategy/trend"
)

func TestOrStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/repository/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/or.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	or := strategy.NewOrStrategy("Or Strategy")
	or.Strategies = append(or.Strategies, strategy.NewBuyAndHoldStrategy())
	or.Strategies = append(or.Strategies, trend.NewMacdStrategy())

	actions, outcomes := strategy.ComputeWithOutcome(or, snapshots)
	outcomes = helper.RoundDigits(outcomes, 2)

	err = strategy.CheckResults(results, actions, outcomes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/repository/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	or := strategy.NewOrStrategy("Or Strategy")
	or.Strategies = append(or.Strategies, strategy.NewBuyAndHoldStrategy())
	or.Strategies = append(or.Strategies, trend.NewMacdStrategy())

	report := or.Report(snapshots)

	fileName := "or.html"
	defer os.Remove(fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
