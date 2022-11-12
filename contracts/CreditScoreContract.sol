// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./hardhat/console.sol";


contract CreditScoreContract {
    address public owner;

    struct Pond {
        string id;
        string fishtype;
    }

    struct Credit {
        uint creditAmount;
        bool status;
    }

    struct Spawning {
        uint totalSpawning; 
        string spawningDate;
    }


    mapping(string => Credit[]) public creditTransactions;
    mapping(string => Pond[]) public pondsInformation;
    mapping(string => Spawning[]) public spawningsInformation;
    uint[] public spawningAverage;


    constructor() {
        console.log("Owner contract deployed by:", msg.sender);
        owner = msg.sender; // 'msg.sender' is sender of  current call, contract deployer for a constructo
    }


    function setCredit(string memory _nik,Credit memory _credit) public {
        creditTransactions[_nik].push(_credit);
    }

    function getCredit(string memory _nik) public view returns (uint){
        uint total = 0;
        for (uint i = 0; i < creditTransactions[_nik].length; i++) {
            total += creditTransactions[_nik][i].creditAmount;
        }
        return total;
    }

    function setPond(string memory _nik,Pond memory _pond) public {
        pondsInformation[_nik].push(_pond);
    }

    function getPond(string memory _nik) public view returns (Pond[] memory){
        return pondsInformation[_nik];
    }

    function setSpawning(string memory _nik,Spawning memory _spawning) public {
        spawningsInformation[_nik].push(_spawning);
    }

    function getSpawning(string memory _nik) public view returns (Spawning[] memory){
        return spawningsInformation[_nik];
    }

    function addSpawningAverage(uint _spawningaverage) public {
        spawningAverage.push(_spawningaverage);
    }

    function getSpawningCurrentAverage() public view returns (uint) {
        return spawningAverage[spawningAverage.length-1];
    }

    function creditValidation(Credit[] memory _credits) public pure returns (bool){
        bool isFeasible = true;
        for (uint i = 0; i < _credits.length; i++) {
            if (!_credits[i].status){
                isFeasible = !isFeasible;
            }
        }
        return isFeasible;
    }

    function pondValidation(uint _pondlength) public pure returns (bool){
        if (_pondlength < 2) {
            return false;
        } else {
            return true;
        }
    }

    function spawningValidation(Spawning[] memory _spawnings) public view returns (bool){
        uint total = 0;
        for (uint i = 0; i < _spawnings.length; i++) {
           total += _spawnings[i].totalSpawning;
        }

        total = total / _spawnings.length;

        if (total >= spawningAverage[spawningAverage.length-1]){
            return true;
        } 

        if (total < spawningAverage[spawningAverage.length-1]){
            return false;
        }

        return true;
       
    }

}

