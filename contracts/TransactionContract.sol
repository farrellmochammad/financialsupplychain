// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./hardhat/console.sol";

contract TransactionContract {
    address public owner;

    enum Status {
        Draft,
        InReview,
        Rejected,
        Approved
    }    

    struct ApproverTransaction {
        string reviewer;
        Status status;
        string filepath;
        string timestamp;
    }

    struct FunderApproverTransaction {
        string reviewer;
        Status status;
        string filepath;
        string timestamp;
        uint modalgiven;
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

    function getApproverValidation(string memory _nik) public view returns (bool){
        bool isApproved = false;
        for (uint i = 0; i < approverTransactions[_nik].length; i++) {
            if (approverTransactions[_nik][i].status == Status.Approved){
                isApproved = true;
            }
        }
        return isApproved;
    }

    function setFunderApproverTransaction(string memory _nik,FunderApproverTransaction memory _funderapprovertransaction) public {
        require(msg.sender == owner);
        require(getApproverValidation(_nik));
        funderApproverTransactions[_nik].push(_funderapprovertransaction);
    }

    function getFunderApproverTransaction(string memory _nik) public view returns (FunderApproverTransaction[] memory){
        return funderApproverTransactions[_nik];
    }

    function getFunderApproverValidation(string memory _nik) public view returns (bool){
        bool isApproved = false;
        for (uint i = 0; i < funderApproverTransactions[_nik].length; i++) {
            if (funderApproverTransactions[_nik][i].status == Status.Approved){
                isApproved = true;
            }
        }
        return isApproved;
    }

    function setMonitoring(string memory _nik,Monitoring memory _monitoring) public {
        require(msg.sender == owner);
        require(getFunderApproverValidation(_nik));
        pondMonitorings[_nik].push(_monitoring);
    }

    function getMonitoring(string memory _nik) public view returns (Monitoring[] memory){
        return pondMonitorings[_nik];
    }

    function setSpawning(string memory _nik,Spawning memory _spawning) public {
        require(msg.sender == owner);
        require(getFunderApproverValidation(_nik));
        spawningHistory[_nik].push(_spawning);
    }

    function getSpawning(string memory _nik) public view returns (Spawning[] memory){
        return spawningHistory[_nik];
    }
   
}