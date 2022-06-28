import React from "react";
import './MemberFind.css';

const MemberFind = () => {

    const users = [
        {
            FirstName:   "วิจิตร",
            LastName:    "ศรีสอ้าน",
            Num:         "0001",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "กิตติมศักดิ์",
        },
        {
            FirstName:   "วิศิษฎ์พร",
            LastName:    "วัฒนวาทิน",
            Num:         "0002",
		    Branch:      "สำนักวิชาเทคโนโลยีสังคม",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "เอมอร",
            LastName:    "ทัศนศร",
            Num:         "0003",
		    Branch:      "สำนักวิชาวิศวกรรมศาสตร์",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "หนึ่ง",
            LastName:    "เตียอำรุง",
            Num:         "0004",
		    Branch:      "สำนักวิชาเทคโนโลยีการเกษตร",
            Email:       "neung@sut.ac.th",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "ไทย",
            LastName:    "ทิพย์สุวรรณกุล",
            Num:         "0005",
		    Branch:      "สำนักวิชาเทคโนโลยีการเกษตร",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "ม.วลัยลักษณ์",
            SubDistrict: "-",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "อรชุน",
            LastName:    "ไชยเสนะ",
            Num:         "0006",
		    Branch:      "สำนักวิชาวิทยาศาสตร์",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "จุฑามาศ",
            LastName:    "สวัสดี",
            Num:         "0007",
		    Branch:      "-",
            Email:       "chuthas@sut.ac.th",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "มนู",
            LastName:    "อัศวเสนา",
            Num:         "0008",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "-",
            SubDistrict: "-",
            District:    "-",
            Province:    "-",
            Category:    "สามัญ",
        },
        {
            FirstName:   "ตริตาภรณ์",
            LastName:    "ชูศรี",
            Num:         "0009",
		    Branch:      "สำนักวิชาวิทยาศาสตร์",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "ธวัชชัย",
            LastName:    "ภิญโญรัฐกานต์",
            Num:         "0010",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "-",
            SubDistrict: "-",
            District:    "-",
            Province:    "-",
            Category:    "สามัญ",
        },
        {
            FirstName:   "วุฒิ",
            LastName:    "ด่านกิตติกุล",
            Num:         "0011",
		    Branch:      "สำนักวิชาวิศวกรรมศาสตร์",
            Email:       "wut@sut.ac.th",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "อุทัย",
            LastName:    "มีคำ",
            Num:         "0012",
		    Branch:      "สำนักวิชาวิศวกรรมศาสตร์",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "กนต์ธร",
            LastName:    "ชำนิประศาสน์",
            Num:         "0013",
		    Branch:      "สำนักวิชาวิศวกรรมศาสตร์",
            Email:       "kontorn@sut.ac.th",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "ราชัย",
            LastName:    "อัศเวศน์",
            Num:         "0014",
		    Branch:      "-",
            Email:       "rachai@sut.ac.th",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "เบญจมาศ",
            LastName:    "จิตรสมบูรณ์",
            Num:         "0015",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "-",
            SubDistrict: "-",
            District:    "-",
            Province:    "-",
            Category:    "สามัญ",
        },
        {
            FirstName:   "อาภรณ์",
            LastName:    "ปราบริปูตลุง",
            Num:         "0016",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "-",
            SubDistrict: "-",
            District:    "-",
            Province:    "-",
            Category:    "สามัญ",
        },
        {
            FirstName:   "Joewono",
            LastName:    "Widjaja",
            Num:         "0017",
		    Branch:      "สำนักวิชาวิทยาศาสตร์",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "เกษม",
            LastName:    "ปราบริปูตลุง",
            Num:         "0018",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "-",
            SubDistrict: "-",
            District:    "-",
            Province:    "-",
            Category:    "สามัญ",
        },
        {
            FirstName:   "ณรงค์",
            LastName:    "อัครพัฒนากูล",
            Num:         "0019",
		    Branch:      "สำนักวิชาวิศวกรรมศาสตร์",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "สามัญ",
        },
        {
            FirstName:   "ธารา",
            LastName:    "เล็กอุทัย",
            Num:         "0020",
		    Branch:      "-",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "-",
            SubDistrict: "-",
            District:    "-",
            Province:    "-",
            Category:    "สามัญ",
        },
        {
            FirstName:   "วิจิตร",
            LastName:    "สอ้าน",
            Num:         "0021",
		    Branch:      "วิศวกรรม",
            Email:       "-",
		    Tel:         "0801234567",
		    Address:     "111 30000",
            SubDistrict: "สุรนารี",
            District:    "เมือง",
            Province:    "นครราชสีมา",
            Category:    "กิตติมศักดิ์",
        },
    ]

    const [userList, setUserList] = React.useState<
        { FirstName: string, LastName: string, Num: string, Branch: string, Email: string, Tel: string, Address: string, SubDistrict: string, District: string, Province: string, Category: string }[] | undefined >(users);
    const [text, setText] = React.useState<string>('');

    const handleOnClick = () => {
        const findUsers = 
            userList && userList?.length > 0 ? userList?.filter(u => u?.FirstName === text)
            : undefined;

            console.log(findUsers);
        setUserList(findUsers);
    }
    return(
        <div>
            <div className="title">
                <h1>ค้นหาสมาชิกสมาคมฯ</h1>
            </div>
            <div className="input_wrapper">
                <input 
                  type="text" 
                  placeholder="ชื่อ" 
                  value={text} 
                  onChange={(e) => {
                    setText(e.target.value); 
                    setUserList(users);
                  }}
                />
                <button disabled={!text} onClick={handleOnClick}>
                    ค้นหา
                </button>
            </div>

            <div className="body">

{userList && userList?.length === 0 && (
    <div className="notFound">ไม่พบ</div>
)}

                {userList && 
                userList?.length> 0 && 
                userList?.map((user) => {
                    return(
                        <div className="body__item">
                            <h3>ชื่อ: {user?.FirstName}</h3>
                            <p>นามสกุล: {user?.LastName}</p>
                            <p>เลขที่สมาชิก: {user?.Num}</p>
                            <p>สาขา: {user?.Branch}</p>
                            <p>อีเมล: {user?.Email}</p>
                            <p>เบอร์โทร: {user?.Tel}</p>
                            <p>ที่อยู่: {user?.Address}</p>
                            <p>ตำบล: {user?.SubDistrict}</p>
                            <p>อำเภอ: {user?.District}</p>
                            <p>จังหวัด: {user?.Province}</p>
                            <p>ประเภทสมาชิก: {user?.Category}</p>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}

export default MemberFind;