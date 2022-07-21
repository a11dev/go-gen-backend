package data

import "github.com/a11dev/go-gen-backend/internal/model"

var TaskList = []model.Task{
	{Key: 978897386, Number: 707427, Status: "Assigned", Typecategory: "AA", Refid: "REFID_000001"},
	{Key: 978897369, Number: 707426, Status: "Working", Typecategory: "AB", Refid: "REFID_000002"},
	{Key: 978897180, Number: 707406, Status: "Working", Typecategory: "AB", Refid: "REFID_000003"},
	{Key: 978897047, Number: 707389, Status: "Deleted", Typecategory: "AA", Refid: "REFID_000004"},
	{Key: 978897038, Number: 707388, Status: "Deleted", Typecategory: "AB", Refid: "REFID_000005"},
	{Key: 978896992, Number: 707387, Status: "Deleted", Typecategory: "AC", Refid: "REFID_000006"},
	{Key: 978896993, Number: 707386, Status: "Deleted", Typecategory: "AA", Refid: "REFID_000007"},
	{Key: 978888483, Number: 707378, Status: "Deleted", Typecategory: "AC", Refid: "REFID_000008"},
	{Key: 978888466, Number: 707377, Status: "Assigned", Typecategory: "AC", Refid: "REFID_000009"},
	{Key: 978888413, Number: 707376, Status: "Working", Typecategory: "AA", Refid: "REFID_000010"},
	{Key: 978886577, Number: 707374, Status: "Assigned", Typecategory: "AC", Refid: "REFID_000011"},
	{Key: 978886576, Number: 707373, Status: "Assigned", Typecategory: "AC", Refid: "REFID_000012"},
	{Key: 978886571, Number: 707372, Status: "Assigned", Typecategory: "AA", Refid: "REFID_000013"},
	{Key: 978886570, Number: 707371, Status: "Assigned", Typecategory: "AB", Refid: "REFID_000014"},
}

var TaskDetails = []model.TaskDetail{
	{Key: 978897386, Body: []byte(`{"Key":978897386, "Number":707427, "Status":"Assigned", "Tpecategory":"AA", "Refid":"REFID_000001"}`)},
	{Key: 978897369, Body: []byte(`{"Key":978897369, "Number":707426, "Status":"Working", "Tpecategory":"AB", "Refid":"REFID_000002"}`)},
	{Key: 978897180, Body: []byte(`{"Key":978897180, "Number":707406, "Status":"Working", "Tpecategory":"AB", "Refid":"REFID_000003"}`)},
	{Key: 978897047, Body: []byte(`{"Key":978897047, "Number":707389, "Status":"Deleted", "Tpecategory":"AA", "Refid":"REFID_000004"}`)},
	{Key: 978897038, Body: []byte(`{"Key":978897038, "Number":707388, "Status":"Deleted", "Tpecategory":"AB", "Refid":"REFID_000005"}`)},
	{Key: 978896992, Body: []byte(`{"Key":978896992, "Number":707387, "Status":"Deleted", "Tpecategory":"AC", "Refid":"REFID_000006"}`)},
	{Key: 978896993, Body: []byte(`{"Key":978896993, "Number":707386, "Status":"Deleted", "Tpecategory":"AA", "Refid":"REFID_000007"}`)},
	{Key: 978888483, Body: []byte(`{"Key":978888483, "Number":707378, "Status":"Deleted", "Tpecategory":"AC", "Refid":"REFID_000008"}`)},
	{Key: 978888466, Body: []byte(`{"Key":978888466, "Number":707377, "Status":"Assigned", "Tpecategory":"AC", "Refid":"REFID_000009"}`)},
	{Key: 978888413, Body: []byte(`{"Key":978888413, "Number":707376, "Status":"Working", "Tpecategory":"AA", "Refid":"REFID_000010"}`)},
	{Key: 978886577, Body: []byte(`{"Key":978886577, "Number":707374, "Status":"Assigned", "Tpecategory":"AC", "Refid":"REFID_000011"}`)},
	{Key: 978886576, Body: []byte(`{"Key":978886576, "Number":707373, "Status":"Assigned", "Tpecategory":"AC", "Refid":"REFID_000012"}`)},
	{Key: 978886571, Body: []byte(`{"Key":978886571, "Number":707372, "Status":"Assigned", "Tpecategory":"AA", "Refid":"REFID_000013"}`)},
	{Key: 978886570, Body: []byte(`{"Key":978886570, "Number":707371, "Status":"Assigned", "Tpecategory":"AB", "Refid":"REFID_000014"}`)},
}
