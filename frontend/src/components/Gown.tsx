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
import { GownInterface } from "../models/IGown";
import { format } from 'date-fns'

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

function Gown() {
  const classes = useStyles();
  const [gown, setGown] = useState<GownInterface[]>([]);
  
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" 
    },
  };

  
  const getGown = async () => {
    fetch(`${apiUrl}/gown`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setGown(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getGown();
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
              ระบบชุดครุย
            </Typography>
          </Box>
          <Box >
            <Button
              component={RouterLink}
              to="/gown/create"
              variant="contained"
              color="primary"
            >
              เพิ่มรายการ
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
                <TableCell align="center" width="20%">
                  ชื่อ-นามสกุล
                </TableCell>
                <TableCell align="center" width="10%">
                  รหัสนักศึกษา
                </TableCell>    
                <TableCell align="center" width="10%">
                  เช่า-ซื้อ-คืน
                </TableCell> 
                <TableCell align="center" width="10%">
                  ชุดครุยไซส์
                </TableCell>     
                <TableCell align="center" width="10%">
                  เบลเซอร์
                </TableCell>       
                <TableCell align="center" width="10%">
                  ระดับปริญญา
                </TableCell>
                <TableCell align="center" width="20%">
                  สาขา
                </TableCell>
                <TableCell align="center" width="10%">
                  เบอร์โทร
                </TableCell>
                <TableCell align="center" width="10%">
                  วันที่
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {gown.map((item: GownInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Name}</TableCell>
                  <TableCell align="center">{item.Student}</TableCell>
                  <TableCell align="center">{item.Status}</TableCell>
                  <TableCell align="center">{item.Size.Size}</TableCell>
                  <TableCell align="center">{item.Blazer.Blazer}</TableCell>
                  <TableCell align="center">{item.Degree.Degree}</TableCell>
                  <TableCell align="center">{item.Branch.Branch}</TableCell>
                  <TableCell align="center">{item.Tel}</TableCell>
                  <TableCell align="center">{format((new Date(item.Date)), 'dd MM yyyy')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Gown;
