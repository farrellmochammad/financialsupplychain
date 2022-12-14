pragma solidity ^0.8.0;

contract ApproverContract {
    string approverid;
    string pondid;
    string nik;
    string approvetimestamp;
    enum Status {
        Draft,
        InReview,
        Rejected,
        Approved
    }    
    Status public status;

    struct ApproverTransaction {
        string from;
        string to;
        Status status;
        string filepath;
        string timestamp;
    }
    ApproverTransaction[] public approvertransactions;

    address public admin;

    constructor(string memory _approverid,string memory _pondid,string memory _nik) {
        admin = msg.sender;
        approverid = _approverid;
        pondid = _pondid;
        nik = _nik;
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

    function addApprovalTransaction(string memory _from,string memory _to, Status _status,string memory _filepath, string memory _timestamp) public {
        approvertransactions.push(
            ApproverTransaction({
                from: _from,
                to: _to,
                status: _status,
                filepath: _filepath,
                timestamp: _timestamp 
            })
        );
    }

    function getApprovalTransactions() public view returns (ApproverTransaction[] memory){
        return approvertransactions;
    }
}