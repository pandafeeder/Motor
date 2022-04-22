
let Tree = {"roots": [{"sourcefile": "./actions/setup/GetDef.sh", "inputs": [], "outputs": ["def.gz"], "level": 0, "parents": [], "children": ["./actions/floorplan/PlaceMacro.sh"]}, {"sourcefile": "./actions/setup/GetNetlist.sh", "inputs": [], "outputs": ["v.gz"], "level": 0, "parents": [], "children": ["./actions/floorplan/PlaceMacro.sh"]}, {"sourcefile": "./actions/setup/GetSDC.sh", "inputs": [], "outputs": ["sdc.gz"], "level": 0, "parents": [], "children": ["./actions/timing/TimeCts.sh", "./actions/timing/TimeRoute.sh", "./actions/timing/TimePlace.sh", "./actions/place/Place.sh"]}], "nodes": [{"sourcefile": "./actions/drc/RouteDRC.sh", "inputs": ["Route.def.gz"], "outputs": ["RouteDRC.rpt"], "level": 6, "parents": ["./actions/route/Route.sh"], "children": []}, {"sourcefile": "./actions/drc/FpDRC.sh", "inputs": ["FinishFp.def.gz"], "outputs": ["FpDRC.rpt"], "level": 3, "parents": ["./actions/floorplan/FinishFp.sh"], "children": []}, {"sourcefile": "./actions/route/Route.sh", "inputs": ["Cts.def.gz", "Cts.v.gz"], "outputs": ["Route.def.gz", "Route.v.gz"], "level": 5, "parents": ["./actions/cts/Cts.sh"], "children": ["./actions/drc/RouteDRC.sh", "./actions/timing/TimeRoute.sh"]}, {"sourcefile": "./actions/cts/Cts.sh", "inputs": ["Place.def.gz", "Place.v.gz"], "outputs": ["Cts.def.gz", "Cts.v.gz"], "level": 4, "parents": ["./actions/place/Place.sh"], "children": ["./actions/route/Route.sh", "./actions/timing/TimeCts.sh"]}, {"sourcefile": "./actions/setup/GetDef.sh", "inputs": [], "outputs": ["def.gz"], "level": 0, "parents": [], "children": ["./actions/floorplan/PlaceMacro.sh"]}, {"sourcefile": "./actions/setup/GetNetlist.sh", "inputs": [], "outputs": ["v.gz"], "level": 0, "parents": [], "children": ["./actions/floorplan/PlaceMacro.sh"]}, {"sourcefile": "./actions/setup/GetSDC.sh", "inputs": [], "outputs": ["sdc.gz"], "level": 0, "parents": [], "children": ["./actions/timing/TimeCts.sh", "./actions/timing/TimeRoute.sh", "./actions/timing/TimePlace.sh", "./actions/place/Place.sh"]}, {"sourcefile": "./actions/timing/TimeCts.sh", "inputs": ["Cts.def.gz", "Cts.v.gz", "sdc.gz"], "outputs": ["TimeCts.rpt"], "level": 5, "parents": ["./actions/cts/Cts.sh", "./actions/setup/GetSDC.sh"], "children": []}, {"sourcefile": "./actions/timing/TimeRoute.sh", "inputs": ["Route.def.gz", "Route.v.gz", "sdc.gz"], "outputs": ["TimeRoute.rpt"], "level": 6, "parents": ["./actions/route/Route.sh", "./actions/setup/GetSDC.sh"], "children": []}, {"sourcefile": "./actions/timing/TimePlace.sh", "inputs": ["Place.v.gz", "Place.def.gz", "sdc.gz"], "outputs": ["TimePlace.rpt"], "level": 4, "parents": ["./actions/place/Place.sh", "./actions/setup/GetSDC.sh"], "children": []}, {"sourcefile": "./actions/floorplan/PlaceMacro.sh", "inputs": ["def.gz", "v.gz"], "outputs": ["PlaceMacro.def.gz", "PlaceMacro.v.gz"], "level": 1, "parents": ["./actions/setup/GetDef.sh", "./actions/setup/GetNetlist.sh"], "children": ["./actions/floorplan/FinishFp.sh"]}, {"sourcefile": "./actions/floorplan/FinishFp.sh", "inputs": ["PlaceMacro.def.gz", "PlaceMacro.v.gz"], "outputs": ["FinishFp.def.gz", "FinishFp.v.gz"], "level": 2, "parents": ["./actions/floorplan/PlaceMacro.sh"], "children": ["./actions/drc/FpDRC.sh", "./actions/place/Place.sh"]}, {"sourcefile": "./actions/place/Place.sh", "inputs": ["FinishFp.def.gz", "FinishFp.v.gz", "sdc.gz"], "outputs": ["Place.def.gz", "Place.v.gz"], "level": 3, "parents": ["./actions/floorplan/FinishFp.sh", "./actions/setup/GetSDC.sh"], "children": ["./actions/cts/Cts.sh", "./actions/timing/TimePlace.sh"]}], "depth": 7}



export default Tree
