// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./hardhat/console.sol";

contract TransactionContract {
    address public owner;


    struct ApproverTransaction {
        string nik;
        string submitby;
        string status;
        string timestamp;
        uint numofponds;
        uint spawningaverage;
        uint creditscore;
    }

    struct FunderApproverTransaction {
        string fundid;
        string nik;
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

    ApproverTransaction[] public approverTransactions;
    FunderApproverTransaction[] public funderApproverTransactions;
    mapping(string => Monitoring[]) public pondMonitorings;
    mapping(string => Spawning[]) public spawningHistory;

     constructor() {
        console.log("Owner contract deployed by:", msg.sender);
        owner = msg.sender; // 'msg.sender' is sender of  current call, contract deployer for a constructo
    }

    function setApproverTransaction(ApproverTransaction memory _approvertransaction) public {
        require(msg.sender == owner);
        approverTransactions.push(_approvertransaction);
    }

    function getApproverTransaction() public view returns (ApproverTransaction[] memory){
        return approverTransactions;
    }


    function setFunderApproverTransaction(FunderApproverTransaction memory _funderapprovertransaction) public {
        require(msg.sender == owner);
        funderApproverTransactions.push(_funderapprovertransaction);
    }

    function getFunderApproverTransaction() public view returns (FunderApproverTransaction[] memory){
        return funderApproverTransactions;
    }

    function setMonitoring(string memory _fundid,Monitoring memory _monitoring) public {
        require(msg.sender == owner);
        pondMonitorings[_fundid].push(_monitoring);
    }

    function getMonitoring(string memory _fundid) public view returns (Monitoring[] memory){
        return pondMonitorings[_fundid];
    }

    function setSpawning(string memory _fundid,Spawning memory _spawning) public {
        require(msg.sender == owner);
        spawningHistory[_fundid].push(_spawning);
    }

    function getSpawning(string memory _fundid) public view returns (Spawning[] memory){
        return spawningHistory[_fundid];
    }
   
}