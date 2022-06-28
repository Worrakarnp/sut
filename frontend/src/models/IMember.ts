import { BranchInterface } from "./IBranch";
import { CategoryInterface } from "./ICategory";
import { SubDistrictInterface } from "./ISubDistrict";
import { DistrictInterface } from "./IDistrict";
import { ProvinceInterface } from "./IProvince";

export interface MemberInterface {
    ID: number,
    Name: string,
    Num: string,
    BranchID: number,
    Branch: BranchInterface,
    Email: string,
    Tel: string,
    Address: string,
	SubDistrictID: number,
    SubDistrict:   SubDistrictInterface,
    DistrictID: number,
    District: DistrictInterface,
    ProvinceID: number,
    Province: ProvinceInterface,
    CategoryID: number,
	Category:   CategoryInterface,
    Date: Date,
}   
