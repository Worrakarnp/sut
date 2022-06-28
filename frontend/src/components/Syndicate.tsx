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
import { SyndicateInterface } from "../models/ISyndicate";

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

function Syndicate() {
  const classes = useStyles();
  const [syndicate, setSyndicate] = useState<SyndicateInterface[]>([]);
  
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" 
    },
  };

  const getSyndicate = async () => {
    fetch(`${apiUrl}/syndicate`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setSyndicate(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getSyndicate();
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
              คณะกรรมการสมาคมเทคโนโลยีสุรนารี ประจำปี 2564-2566
            </Typography>
          </Box>
          <Box >
            <Button
              component={RouterLink}
              to="/syndicate/create"
              variant="contained"
              color="primary"
            >
              เพิ่มชื่อคณะกรรมการ
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
                <TableCell align="center" width="10%">
                  ชื่อ-นามสกุล
                </TableCell>
                <TableCell align="center" width="10%">
                  ตำแหน่ง
                </TableCell>
                <TableCell align="center" width="10%">
                  สาขา
                </TableCell>
                <TableCell align="center" width="10%">
                  เบอร์โทร
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {syndicate.map((item: SyndicateInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Name}</TableCell>
                  <TableCell align="center">{item.Rank}</TableCell>
                  <TableCell align="center">{item.Branch.Branch}</TableCell>
                  <TableCell align="center">{item.Tel}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Syndicate;
