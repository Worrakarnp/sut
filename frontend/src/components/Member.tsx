import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { format } from 'date-fns'
import { MemberInterface } from "../models/IMember";

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

function Member() {
  const classes = useStyles();
  const [member, setMember] = useState<MemberInterface[]>([]);  
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" 
    },
  };
 
  const getMember = async () => {
    fetch(`${apiUrl}/member`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setMember(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getMember();
  }, []);  
  
  return (
    <div>
      <Container className={classes.container} maxWidth="lg">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              สมาชิกสมาคมเทคโนโลยีสุรนารี
            </Typography>
          </Box>
          <Box >
            <Button
              component={RouterLink}
              to="/member/create"
              variant="contained"
              color="primary"
            >
              เพิ่มสมาชิก
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="2%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="50%">
                  ชื่อ-นามสกุล
                </TableCell>
                <TableCell align="center" width="10%">
                  เลขที่สมาชิก
                </TableCell>                
                <TableCell align="center" width="40%">
                  สาขา
                </TableCell>
                <TableCell align="center" width="20%">
                  อีเมล
                </TableCell>
                <TableCell align="center" width="10%">
                  เบอร์โทร
                </TableCell>
                <TableCell align="center" width="20%">
                  ที่อยู่
                </TableCell>
                <TableCell align="center" width="10%">
                  ตำบล
                </TableCell>
                <TableCell align="center" width="10%">
                  อำเภอ
                </TableCell>
                <TableCell align="center" width="10%">
                  จังหวัด
                </TableCell>
                <TableCell align="center" width="10%">
                  วันที่
                </TableCell>
                <TableCell align="center" width="5%">
                  ประเภทสมาชิก
               </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {member.map((item: MemberInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Name}</TableCell>
                  <TableCell align="center">{item.Num}</TableCell>
                  <TableCell align="center">{item.Branch.Branch}</TableCell>
                  <TableCell align="center">{item.Email}</TableCell>
                  <TableCell align="center">{item.Tel}</TableCell>
                  <TableCell align="center">{item.Address}</TableCell>
                  <TableCell align="center">{item.SubDistrict.SubDistrict}</TableCell>
                  <TableCell align="center">{item.District.District}</TableCell>
                  <TableCell align="center">{item.Province.Province}</TableCell>
                  <TableCell align="center">{format((new Date(item.Date)), 'dd MM yyyy')}</TableCell>
                  <TableCell align="center">{item.Category.Category}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Member;
