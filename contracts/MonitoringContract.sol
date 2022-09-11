pragma solidity ^0.8.0;

contract MonitoringContract {
    string monitoringid;
    string pondid;
    string pondname;
    string area;
    string district;
    string fishtype;

    
    struct Monitoring {
        string timestamp;
        uint weight;
        uint temperature;
        uint humidity;
    }

    Monitoring[] public monitorings;

    struct Spawning {
        string timestamp;
        uint weight;
    }

    Spawning[] public spawnings;

    address public admin;

    constructor(string memory _monitoringid,string memory _pondid, string memory _pondname, string memory _area, string memory _district, string memory _fishtype) {
        admin = msg.sender;
        monitoringid = _monitoringid;
        pondid = _pondid;
        pondname = _pondname;
        area = _area;
        district = _district;
        fishtype = _fishtype;
    }

    function addMonitoring(string memory _timestamp, uint _weight, uint _temperature, uint _humidity) public {
        monitorings.push(
            Monitoring({
                timestamp: _timestamp,
                weight: _weight,
                temperature: _temperature,
                humidity: _humidity
            })
        );
    }

    function getMonitorings() public view returns (Monitoring[] memory){
        return monitorings;
    }

    function addSpawning(string memory _timestamp, uint _weight) public {
        spawnings.push(
            Spawning({
                timestamp: _timestamp,
                weight: _weight
            })
        );
    }

    function getSpawnings() public view returns (Spawning[] memory){
        return spawnings;
    }


  
}