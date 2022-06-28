import React from "react";
import './UserFind.css';

const UserFind = () => {

    const users = [
        {
            FirstName:  "รัฐมงคล",
            LastName:   "เหมรัตนานนท์",
		    Rank:       "นายกสมาคม",
		    Branch:     "วิศวกรรมโลหการ",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "ศิริชัย",
            LastName:   "นามวิเศษ",
		    Rank:       "อุปนายก-กิจกรรม",
		    Branch:     "วิศวกรรมไฟฟ้า",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "กฤดิภัค",
            LastName:   "โกยดุลย์",
            Rank:       "กรรมการ-กิจกรรม",
            Branch:     "วิศวกรรมเซรามิก",
            Tel:        "0801234567",
        },
        {
            FirstName:  "กีรติคุณ",
            LastName:   "บุญประสพ",
		    Rank:       "กรรมการ-กิจกรรม",
		    Branch:     "วิศวกรรมไฟฟ้า",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "อรุณี",
            LastName:   "อะโสต",
		    Rank:       "กรรมการ-กิจกรรม",
		    Branch:     "วิศวกรรมเครื่องกล",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "ทรรศสิน",
            LastName:   "เดชะกูล",
		    Rank:       "กรรมการ-กิจกรรม",
		    Branch:     "วิศวกรรมขนส่ง",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "อมฤต",
            LastName:   "ฉายะรถี",
		    Rank:       "กรรมการ-กิจกรรม",
		    Branch:     "วิศวกรรมอุตสาหการ",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "สมชาย",
            LastName:   "สุขอินทร์",
		    Rank:       "อุปนายก-วิชาการ",
		    Branch:     "คณิตศาสตร์ประยุกต์",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "วชิระ",
            LastName:   "พุทธิแจ่ม",
		    Rank:       "กรรมการ-วิชาการ",
		    Branch:     "วิศวกรรมเครื่องกล",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "ปรัชญา",
            LastName:   "เทพณรงค์",
		    Rank:       "กรรมการ-วิชาการ",
		    Branch:     "วิศวกรรมเทคโนโลยีธรณี",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "เพิ่มพูน",
            LastName:   "เสมอ",
		    Rank:       "กรรมการ-วิชาการ",
		    Branch:     "วิศวกรรมไฟฟ้า",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "กงจักร",
            LastName:   "ลมวิชัย",
		    Rank:       "กรรมการ-วิชาการ",
		    Branch:     "วิศวกรรมเกษตรและอาหาร",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "วิยภรณ์",
            LastName:   "กรองทอง",
		    Rank:       "กรรมการ-วิชาการ",
		    Branch:     "วิศวกรรมเซรามิก",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "วิศรุตน์",
            LastName:   "ไชยเพ็ชร",
		    Rank:       "เลขาธิการ",
		    Branch:     "วิศวกรรมอุตสาหการ",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "บรรพต",
            LastName:   "สุทธิพันธ์",
		    Rank:       "กรรมการ-เลขาธิการ",
		    Branch:     "วิศวกรรมเทคโนโลยีธรณี",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "รัฐพงษ์",
            LastName:   "อ่อนจันทร์",
		    Rank:       "เหรัญญิก",
		    Branch:     "วิศวกรรมคอมพิวเตอร์",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "วีระพันธ์",
            LastName:   "เต็มดวง",
		    Rank:       "กรรมการ-การเงิน",
		    Branch:     "วิศวกรรมอุตสาหการ",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "พันธกานต์",
            LastName:   "รื่นกลิ่น",
		    Rank:       "ปฏิคม",
		    Branch:     "นิเทศศาสตร์",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "คมเดช",
            LastName:   "ธูปกิ่ง",
		    Rank:       "กรรมการ-ปฏิคม",
		    Branch:     "วิศวกรรมเซรามิก",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "ธนาทร",
            LastName:   "อารยวริต",
		    Rank:       "นายทะเบียน",
		    Branch:     "วิศวกรรมโยธาและการบริหารงานก่อสร้าง",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "ฉัฐพร",
            LastName:   "ทองวิจิตต์",
		    Rank:       "กรรมการ-ทะเบียน",
		    Branch:     "วิศวกรรมสิ่งแวดล้อม",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "อดิศักดิ์",
            LastName:   "ยั่งยืน",
		    Rank:       "ประชาสัมพันธ์",
		    Branch:     "วิศวกรรมสิ่งแวดล้อม",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "เอกชัย",
            LastName:   "บุญรสศักดิ์",
		    Rank:       "กรรมการ-ประชาสัมพันธ์",
		    Branch:     "วิศวกรรมคอมพิวเตอร์",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "อุดมศักดิ์",
            LastName:   "บัวสำราญ",
		    Rank:       "กรรมการ-ประชาสัมพันธ์",
		    Branch:     "เทคโนโลยีอาหาร",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "สุเมธี",
            LastName:   "ถีสูงเนิน",
		    Rank:       "กรรมการ-ประชาสัมพันธ์",
		    Branch:     "วิศวกรรมแมคคาทรอนิกส์",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "จิรนันท์",
            LastName:   "มหาดไทย",
		    Rank:       "เจ้าหน้าที่บริหารทั่วไป",
		    Branch:     "เทคโนโลยีผลิตพืช",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "สุชาดา",
            LastName:   "สงวนพันธุ์",
		    Rank:       "เจ้าหน้าที่บริหารทั่วไป",
		    Branch:     "การจัดการการตลาด",
		    Tel:        "0801234567",
        },
        {
            FirstName:  "วิจิตรา",
            LastName:   "แสนโคตร",
		    Rank:       "เจ้าหน้าที่บริหารทั่วไป",
		    Branch:     "การจัดการโลจิสติกส์",
		    Tel:        "0801234567",
        },
    ]

    const [userList, setUserList] = React.useState<
        { FirstName: string, LastName: string, Rank: string, Branch: string, Tel: string }[] | undefined >(users);
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
                <h1>ค้นหาคณะกรรมการสมาคมฯ</h1>
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
                            <p>ตำแหน่ง: {user?.Rank}</p>
                            <p>สาขา: {user?.Branch}</p>
                            <p>เบอร์โทร: {user?.Tel}</p>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}

export default UserFind;