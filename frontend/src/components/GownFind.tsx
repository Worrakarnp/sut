import React from "react";
import './GownFind.css';

const GownFind = () => {

    const users = [
        {
            FirstName: "วรกานต์",
            LastName:  "ปินชัย",
            Student:   "B1111111",
            Status:    "ซื้อ",
            Size:      "M",
            Blazer:    "-",
            Degree:    "ปริญญาตรี",
		    Branch:    "วิศวกรรมคอมพิวเตอร์",
		    Tel:       "0801234567",
        },
    ]

    const [userList, setUserList] = React.useState<
        { FirstName: string, LastName: string, Student: string, Status: string, Size: string, Blazer: string, Degree: string, Branch: string, Tel: string }[] | undefined >(users);
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
                <h1>ค้นหาชุดครุย</h1>
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
                            <p>รหัสนักศึกษา: {user?.Student}</p>
                            <p>เช่า-ซื้อ-คืน: {user?.Status}</p>
                            <p>ไซส์: {user?.Size}</p>
                            <p>เบลเซอร์: {user?.Blazer}</p>
                            <p>ระดับปริญญา: {user?.Degree}</p>
                            <p>สาขา: {user?.Branch}</p>
                            <p>เบอร์โทร: {user?.Tel}</p>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}

export default GownFind;