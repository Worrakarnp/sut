import React from "react";
import './TraderFind.css';

const TraderFind = () => {

    const users = [
        {
            FirstName:   "วิจิตร",
            LastName:    "ศรีสอ้าน",
		    Tel:     "0801234567",
		    Address: "111 ต.สุรนารี อ.เมือง จ.นครราชสีมา 30000",
        },
    ]

    const [userList, setUserList] = React.useState<
        { FirstName: string, LastName: string, Tel: string, Address: string }[] | undefined >(users);
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
                <h1>ค้นหาผู้ประกอบการ</h1>
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
                            <p>เบอร์โทร: {user?.Tel}</p>
                            <p>ที่อยู่: {user?.Address}</p>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}

export default TraderFind;