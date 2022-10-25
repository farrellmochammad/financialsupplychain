pragma solidity ^0.8.0;

contract CreditContract {

    struct Credit {
        uint creditAmount;
        bool status;
    }

    mapping(string => Credit[]) public creditTransactions;
    mapping(string => bool) public creditTransactionsInserted;


    function setCredit(string memory _nik,Credit memory _credit) public {
        creditTransactions[_nik].push(_credit);
        creditTransactionsInserted[_nik] = true;
    }

    function getCredit(string memory _nik) public view returns (uint){
        require(!creditTransactionsInserted[_nik], "nik not found");
        uint total = 0;
        for (uint i = 0; i < creditTransactions[_nik].length; i++) {
            total += creditTransactions[_nik][i].creditAmount;
        }

        return total;
    }

}

