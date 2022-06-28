import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import HomeIcon from "@material-ui/icons/Home";
import SearchIcon from "@material-ui/icons/Search";
import FaceIcon from '@material-ui/icons/Face';
import GroupIcon from '@material-ui/icons/Group';
import BusinessCenterIcon from '@material-ui/icons/BusinessCenter';
import SchoolIcon from '@material-ui/icons/School';
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ฐานข้อมูลภายในสมาคมเทคโนโลยีสุรนารี</h1>
        <h4>เกี่ยวกับ</h4>
        <p>
        ระบบฐานข้อมูลสมาคม เป็นระบบที่มีไว้สำหรับเจ้าหน้าที่ เป็นระบบที่สามารถบันทึกข้อมูลต่างๆ เกี่ยวกับข้อมูลภายในสมาคมเทคโนโลยีสุรนารี 
        ประกอบไปด้วย 
        </p>
        
        <h4><HomeIcon color="primary" /> หน้าหลัก</h4>
        <h4><SearchIcon color="secondary" /> ค้นหาคณะกรรมการสมาคมฯ</h4>
        <h4><FaceIcon /> คณะกรรมการสมาคมฯ</h4> 
        <h4><SearchIcon color="secondary" /> ค้นหาสมาชิกสมาคมฯ</h4>
        <h4><GroupIcon /> สมาชิกสมาคมฯ</h4>
        <h4><SearchIcon color="secondary" /> ค้นหาผู้ประกอบการ</h4>
        <h4><BusinessCenterIcon /> ผู้ประกอบการ</h4>
        <h4><SearchIcon color="secondary" /> ค้นหาชุดครุย</h4>
        <h4><SchoolIcon /> ระบบชุดครุย</h4>

      </Container>
    </div>
  );
}
export default Home;