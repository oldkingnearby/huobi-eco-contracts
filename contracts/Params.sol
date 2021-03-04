pragma solidity ^0.6.0;

contract Params {
    bool public initialized;

    // System contracts
    address
        public constant ValidatorContractAddr = 0x0000000000000000000000000000000000001111;
    address
        public constant PunishContractAddr = 0x0000000000000000000000000000000000002222;
    address
        public constant ProposalAddr = 0x0000000000000000000000000000000000003333;

    // System params
    uint16 public constant MaxValidators = 21;
    // Validator have to wait StakingLockPeriod blocks to withdraw staking
    uint64 public constant StakingLockPeriod = 86400;
    // Validator have to wait WithdrawProfitPeriod blocks to withdraw his profits
    uint64 public constant WithdrawProfitPeriod = 28800;
    uint256 public constant MinimalStakingCoin = 32 ether;

    modifier onlyMiner() {
        require(msg.sender == block.coinbase, "Miner only");
        _;
    }

    modifier onlyNotInitialized() {
        require(!initialized, "Already initialized");
        _;
    }

    modifier onlyInitialized() {
        require(initialized, "Not init yet");
        _;
    }

    modifier onlyPunishContract() {
        require(msg.sender == PunishContractAddr, "Punish contract only");
        _;
    }

    modifier onlyBlockEpoch(uint256 epoch) {
        require(block.number % epoch == 0, "Block epoch only");
        _;
    }

    modifier onlyValidatorsContract() {
        require(
            msg.sender == ValidatorContractAddr,
            "Validators contract only"
        );
        _;
    }

    modifier onlyProposalContract() {
        require(msg.sender == ProposalAddr, "Proposal contract only");
        _;
    }
}
