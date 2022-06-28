import { SizeInterface } from "./ISize";
import { BlazerInterface } from "./IBlazer";
import { DegreeInterface } from "./IDegree";
import { BranchInterface } from "./IBranch";

export interface GownInterface {
    ID: number,
    Name: string,
    Student: string,
    Status: string,
    SizeID: number,
    Size: SizeInterface,
    BlazerID: number,
    Blazer: BlazerInterface,
    DegreeID: number,
    Degree: DegreeInterface,
    BranchID: number,
    Branch: BranchInterface,
    Tel: string,
    Date: Date,
}   
