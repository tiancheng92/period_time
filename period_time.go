package period_time

import (
	"sort"
	"time"
)

type PeriodTime struct {
	StartTime time.Time
	EndTime   time.Time
}

type PTSlice []PeriodTime

func New() PTSlice {
	var p PTSlice
	p = []PeriodTime{}
	return p
}

func (pt *PTSlice) Append(StartTime, EndTime time.Time) {
	*pt = append(*pt, PeriodTime{
		StartTime: StartTime,
		EndTime:   EndTime,
	})
}

func (pt PTSlice) Len() int { return len(pt) }

func (pt PTSlice) Swap(i, j int) { pt[i], pt[j] = pt[j], pt[i] }

func (pt PTSlice) Less(i, j int) bool { return pt[i].StartTime.Before(pt[j].StartTime) }

func (pt PTSlice) Union() PTSlice {
	var p PTSlice
	if len(pt) > 1 {
		sort.Stable(pt)
		p = append(p, pt[0])
		for k, v := range pt {
			if v.StartTime.Unix() > v.EndTime.Unix() {
				return p
			}
			if k == 0 {
				continue
			}
			if v.StartTime.Unix() >= p[len(p)-1].StartTime.Unix() && v.StartTime.Unix() <= p[len(p)-1].EndTime.Unix() {
				if v.EndTime.Unix() > p[len(p)-1].EndTime.Unix() {
					p[len(p)-1].EndTime = v.EndTime
				}
			} else if v.StartTime.Unix() > p[len(p)-1].EndTime.Unix() {
				inner := PeriodTime{StartTime: v.StartTime, EndTime: v.EndTime}
				p = append(p, inner)
			}
		}
	}
	return p
}
