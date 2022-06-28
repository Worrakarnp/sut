import { BranchInterface } from "./IBranch";

export interface SyndicateInterface {
    ID: number,
    Name: string,
    Rank: string,
    BranchID: number,
    Branch:   BranchInterface,
    Tel: string,
}
   
