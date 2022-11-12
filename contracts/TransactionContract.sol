// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./hardhat/console.sol";

contract TransactionContract {
    address public owner;


    struct ApproverTransaction {
        string submitby;
        string status;
        string timestamp;
        uint numofponds;
        uint spawningaverage;
        uint creditscore;
    }

    struct FunderApproverTransaction {
        string fundid;
        string funder;
        string timestamp;
        uint numberofponds;
        uint amountoffund;
    }

    struct Monitoring {
        string timestamp;
        uint weight;
        uint temperature;
        uint humidity;
    }

    struct Spawning {
        string pondid;
        string timestamp;
        uint weight;
    }

    mapping(string => ApproverTransaction[]) public approverTransactions;
    mapping(string => FunderApproverTransaction[]) public funderApproverTransactions;
    mapping(string => Monitoring[]) public pondMonitorings;
    mapping(string => Spawning[]) public spawningHistory;

     constructor() {
        console.log("Owner contract deployed by:", msg.sender);
        owner = msg.sender; // 'msg.sender' is sender of  current call, contract deployer for a constructo
    }

    function setApproverTransaction(string memory _nik,ApproverTransaction memory _approvertransaction) public {
        require(msg.sender == owner);
        approverTransactions[_nik].push(_approvertransaction);
    }

    function getApproverTransaction(string memory _nik) public view returns (ApproverTransaction[] memory){
        return approverTransactions[_nik];
    }


    function setFunderApproverTransaction(string memory _nik,FunderApproverTransaction memory _funderapprovertransaction) public {
        require(msg.sender == owner);
        funderApproverTransactions[_nik].push(_funderapprovertransaction);
    }

    function getFunderApproverTransaction(string memory _nik) public view returns (FunderApproverTransaction[] memory){
        return funderApproverTransactions[_nik];
    }

    function setMonitoring(string memory _pondid,Monitoring memory _monitoring) public {
        require(msg.sender == owner);
        pondMonitorings[_pondid].push(_monitoring);
    }

    function getMonitoring(string memory _pondid) public view returns (Monitoring[] memory){
        return pondMonitorings[_pondid];
    }

    function setSpawning(string memory _nik,Spawning memory _spawning) public {
        require(msg.sender == owner);
        spawningHistory[_nik].push(_spawning);
    }

    function getSpawning(string memory _nik) public view returns (Spawning[] memory){
        return spawningHistory[_nik];
    }
   
}