pragma solidity ^0.8.0;

contract FunderApproverContract {
    string funderapproverid;
    string agentapproverid;
    string approvetimestamp;
    enum Status {
        Draft,
        InReview,
        Rejected,
        Approved
    }    
    Status public status;

    struct FunderApproverTransaction {
        Status status;
        string filepath;
        string timestamp;
        uint modalgiven;
        uint creditscore;
    }

    FunderApproverTransaction[] public funderapprovertransactions;

    address public admin;

    constructor(string memory _funderapproverid,string memory _agentapproverid) {
        admin = msg.sender;
        funderapproverid = _funderapproverid;
        agentapproverid = _agentapproverid;
        resetStatus();
    }

    // Returns uint
    // Draft  - 0
    // InReview  - 1
    // Rejected - 2
    // Approved - 3
    function getStatus() public view returns (Status) {
        return status;
    }

    // Update status by passing uint into input
    function setStatus(Status _status) public {
        require(msg.sender == admin);
        status = _status;
    }

    function rejectStatus() public {
        require(msg.sender == admin);
        status = Status.Rejected;
    }

    // delete resets the enum to its first value, 0
    function resetStatus() public {
        require(msg.sender == admin);
        delete status;
    }

    function setApproveTimestamp(string memory _timestamp) public {
        require(msg.sender == admin);
        approvetimestamp = _timestamp;
    }

    function addFunderApprovalTransaction(Status _status,string memory _filepath, string memory _timestamp, uint256 _modalgiven, uint256 _creditscore) public {
        funderapprovertransactions.push(
            FunderApproverTransaction({
                status: _status,
                filepath: _filepath,
                timestamp: _timestamp,
                modalgiven: _modalgiven,
                creditscore: _creditscore
            })
        );
    }

    function getFunderApprovalTransactions() public view returns (FunderApproverTransaction[] memory){
        return funderapprovertransactions;
    }
}