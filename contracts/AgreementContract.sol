pragma solidity ^0.8.0;

contract AgreementContract {
    string agreementid;
    address public admin;

    struct AgreementHistory {
        string timestamp;
        string approverfunderid;
        string monitoringid;
        string approverid;
        string signed;
    }

    AgreementHistory[] public agreementhistories;

    constructor(string memory _agreementid) {
        admin = msg.sender;
        agreementid = _agreementid;
    }

    
    function addAgreement(string memory _timestamp, string memory _approverfunderid, string memory _monitoringid, string memory _approverid, string memory _signed) public {
        agreementhistories.push(
            AgreementHistory({
                timestamp: _timestamp,
                approverfunderid: _approverfunderid,
                monitoringid: _monitoringid,
                approverid: _approverid,
                signed: _signed
            })
        );
    }

    function getAgreements() public view returns (AgreementHistory[] memory){
        return agreementhistories;
    }



}