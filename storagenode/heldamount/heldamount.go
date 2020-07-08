// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package heldamount

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"storj.io/common/storj"
)

// DB works with heldamount database
//
// architecture: Database
type DB interface {
	// StorePayStub inserts or updates held amount into the DB
	StorePayStub(ctx context.Context, paystub PayStub) error
	// GetPayStub retrieves paystub stats for specific satellite
	GetPayStub(ctx context.Context, satelliteID storj.NodeID, period string) (*PayStub, error)
	// AllPayStubs retrieves paystub data from all satellites in specific period from DB.
	AllPayStubs(ctx context.Context, period string) ([]PayStub, error)
	// SatellitesHeldbackHistory retrieves heldback history for specific satellite from DB.
	SatellitesHeldbackHistory(ctx context.Context, satelliteID storj.NodeID) ([]AmountPeriod, error)
	// SatellitesDisposedHistory returns all disposed amount for specific satellite from DB.
	SatellitesDisposedHistory(ctx context.Context, satelliteID storj.NodeID) (int64, error)
	// SatellitePeriods retrieves all periods for concrete satellite in which we have some heldamount data.
	SatellitePeriods(ctx context.Context, satelliteID storj.NodeID) ([]string, error)
	// AllPeriods retrieves all periods in which we have some heldamount data.
	AllPeriods(ctx context.Context) ([]string, error)
	// StorePayment inserts or updates payment into the DB
	StorePayment(ctx context.Context, payment Payment) error
}

// ErrNoPayStubForPeriod represents errors from the heldamount database.
var ErrNoPayStubForPeriod = errs.Class("no payStub for period error")

// PayStub is node heldamount data for satellite by specific period.
type PayStub struct {
	SatelliteID    storj.NodeID `json:"satelliteId"`
	Period         string       `json:"period"`
	Created        time.Time    `json:"created"`
	Codes          string       `json:"codes"`
	UsageAtRest    float64      `json:"usageAtRest"`
	UsageGet       int64        `json:"usageGet"`
	UsagePut       int64        `json:"usagePut"`
	UsageGetRepair int64        `json:"usageGetRepair"`
	UsagePutRepair int64        `json:"usagePutRepair"`
	UsageGetAudit  int64        `json:"usageGetAudit"`
	CompAtRest     int64        `json:"compAtRest"`
	CompGet        int64        `json:"compGet"`
	CompPut        int64        `json:"compPut"`
	CompGetRepair  int64        `json:"compGetRepair"`
	CompPutRepair  int64        `json:"compPutRepair"`
	CompGetAudit   int64        `json:"compGetAudit"`
	SurgePercent   int64        `json:"surgePercent"`
	Held           int64        `json:"held"`
	Owed           int64        `json:"owed"`
	Disposed       int64        `json:"disposed"`
	Paid           int64        `json:"paid"`
	Receipt        string       `json:"receipt"`
}

// AmountPeriod is node's held amount for period.
type AmountPeriod struct {
	Period string `json:"period"`
	Held   int64  `json:"held"`
}

// EstimatedPayout contains usage and estimated payout data for current and previous months.
type EstimatedPayout struct {
	CurrentMonth  PayoutMonthly `json:"currentMonth"`
	PreviousMonth PayoutMonthly `json:"previousMonth"`
}

// PayoutMonthly contains usage and estimated payout date.
type PayoutMonthly struct {
	EgressBandwidth         int64   `json:"egressBandwidth"`
	EgressBandwidthPayout   int64   `json:"egressBandwidthPayout"`
	EgressRepairAudit       int64   `json:"egressRepairAudit"`
	EgressRepairAuditPayout int64   `json:"egressRepairAuditPayout"`
	DiskSpace               float64 `json:"diskSpace"`
	DiskSpacePayout         int64   `json:"diskSpacePayout"`
	HeldRate                int64   `json:"heldRate"`
	Payout                  int64   `json:"payout"`
	Held                    int64   `json:"held"`
}

// Payment is node payment data for specific period.
type Payment struct {
	ID          int64        `json:"id"`
	Created     time.Time    `json:"created"`
	SatelliteID storj.NodeID `json:"satelliteId"`
	Period      string       `json:"period"`
	Amount      int64        `json:"amount"`
	Receipt     string       `json:"receipt"`
	Notes       string       `json:"notes"`
}
