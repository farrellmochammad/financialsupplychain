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
        uint256 creditAmount;
        bool status;
    }

    struct Spawning {
        uint256 totalSpawning;
        string spawningDate;
    }

    mapping(string => Credit[]) public creditTransactions;
    mapping(string => Pond[]) public pondsInformation;
    mapping(string => Spawning[]) public spawningsInformation;
    uint256[] public spawningAverage;

    constructor() {
        console.log("Owner contract deployed by:", msg.sender);
        owner = msg.sender; // 'msg.sender' is sender of  current call, contract deployer for a constructo
    }

    function experienceScore(uint256 _yoe) public pure returns (uint256) {
        if (_yoe < 2) {
            return 0;
        }

        if (_yoe >= 2 && _yoe < 5) {
            return 25;
        }

        if (_yoe >= 5 && _yoe < 7) {
            return 38;
        }

        if (_yoe >= 7) {
            return 53;
        }

        return 0;
    }

    function pondScore(uint256 _nop) public pure returns (uint256) {
        if (_nop < 2) {
            return 0;
        }

        if (_nop >= 2 && _nop < 4) {
            return 36;
        }

        if (_nop >= 4 && _nop < 6) {
            return 53;
        }

        if (_nop >= 6 && _nop < 9) {
            return 70;
        }

        if (_nop >= 10) {
            return 94;
        }

        return 0;
    }

    function creditHistoryScore(Credit[] memory _credits)
        public
        pure
        returns (uint256)
    {
        uint256 total = 0;
        for (uint256 i = 0; i < _credits.length; i++) {
            if (!_credits[i].status) {
                total += _credits[i].creditAmount;
            }
        }

        if (total == 0) {
            return 0;
        }

        if (total >= 1 && total < 3) {
            return 15;
        }

        if (total >= 3 && total < 5) {
            return 27;
        }

        if (total >= 5 && total < 8) {
            return 43;
        }

        if (total > 8) {
            return 60;
        }

        return 0;
    }

    function spawningScore(Spawning[] memory _spawnings)
        public
        pure
        returns (uint256)
    {
        uint256 total = 0;
        for (uint256 i = 0; i < _spawnings.length; i++) {
            total += _spawnings[i].totalSpawning;
        }

        if (total < 100) {
            return 15;
        }

        if (total >= 100 && total < 150) {
            return 100;
        }

        if (total >= 150 && total < 200) {
            return 225;
        }

        if (total >= 200) {
            return 350;
        }

        return 0;
    }

    function generateCreditScore(
        uint256 _yoe,
        uint256 _nop,
        Credit[] memory _credits,
        Spawning[] memory _spawnings
    ) public pure returns (uint256) {
        return
            experienceScore(_yoe) +
            pondScore(_nop) +
            creditHistoryScore(_credits) +
            spawningScore(_spawnings);
    }

    function setCredit(string memory _nik, Credit memory _credit) public {
        creditTransactions[_nik].push(_credit);
    }

    function getCredit(string memory _nik) public view returns (uint256) {
        uint256 total = 0;
        for (uint256 i = 0; i < creditTransactions[_nik].length; i++) {
            total += creditTransactions[_nik][i].creditAmount;
        }
        return total;
    }

    function setPond(string memory _nik, Pond memory _pond) public {
        pondsInformation[_nik].push(_pond);
    }

    function getPond(string memory _nik) public view returns (Pond[] memory) {
        return pondsInformation[_nik];
    }

    function setSpawning(string memory _nik, Spawning memory _spawning) public {
        spawningsInformation[_nik].push(_spawning);
    }

    function getSpawning(string memory _nik)
        public
        view
        returns (Spawning[] memory)
    {
        return spawningsInformation[_nik];
    }

    function addSpawningAverage(uint256 _spawningaverage) public {
        spawningAverage.push(_spawningaverage);
    }

    function getSpawningCurrentAverage() public view returns (uint256) {
        return spawningAverage[spawningAverage.length - 1];
    }

    function creditValidation(Credit[] memory _credits)
        public
        pure
        returns (bool)
    {
        bool isFeasible = true;
        for (uint256 i = 0; i < _credits.length; i++) {
            if (!_credits[i].status) {
                isFeasible = !isFeasible;
            }
        }
        return isFeasible;
    }

    function pondValidation(uint256 _pondlength) public pure returns (bool) {
        if (_pondlength < 2) {
            return false;
        } else {
            return true;
        }
    }

    function experienceValidation(uint _durationofexperience)
        public
        pure
        returns (uint)
    {
        if (_durationofexperience < 2) {
            return 0;
        }

        if (_durationofexperience >= 2 && _durationofexperience < 5) {
            return 17;
        }

        if (_durationofexperience >= 5 && _durationofexperience < 8) {
            return 33;
        }

        if (_durationofexperience >= 8) {
            return 45;
        }

        return 0;
    }

    function spawningValidation(Spawning[] memory _spawnings)
        public
        view
        returns (bool)
    {
        uint256 total = 0;
        for (uint256 i = 0; i < _spawnings.length; i++) {
            total += _spawnings[i].totalSpawning;
        }

        total = total / _spawnings.length;

        if (total >= spawningAverage[spawningAverage.length - 1]) {
            return true;
        }

        if (total < spawningAverage[spawningAverage.length - 1]) {
            return false;
        }

        return true;
    }
}
