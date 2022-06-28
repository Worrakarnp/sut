import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { Select } from "@material-ui/core";
import TextField from '@material-ui/core/TextField';

import { BranchInterface } from "../models/IBranch";
import { SyndicateInterface } from "../models/ISyndicate";

function Alert(props: AlertProps) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: { flexGrow: 1 },
        container: { marginTop: theme.spacing(2) },
        paper: { padding: theme.spacing(2), color: theme.palette.text.secondary },
    })
);

function SyndicateCreate() {
  const classes = useStyles();

  const [branches, setBranch] = useState<BranchInterface[]>([]);
  const [Syndicate, setSyndicate] = useState<Partial<SyndicateInterface>>({});

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
   
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof Syndicate;
    setSyndicate({
      ...Syndicate,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof Syndicate;
    const { value } = event.target;
    setSyndicate({ ...Syndicate, [id]: value });
  };

  const getBranch = async () => {
    fetch(`${apiUrl}/branches`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setBranch(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getBranch();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      Name: Syndicate.Name ?? "",
      Rank: Syndicate.Rank ?? "",
      BranchID: convertType(Syndicate.BranchID),
      Tel: Syndicate.Tel,
    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/syndicate`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true)
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true)
          setErrorMessage(res.error)
        }
      });
  }
  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              เพิ่มคณะกรรมการสมาคมฯ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={2} className={classes.root}>
          <Grid item xs={12}>
            <p>ชื่อคณะกรรมการ</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกชื่อคณะกรรมการ"
                value={Syndicate.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  

          <Grid item xs={12}>
            <p>ตำแหน่ง</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Rank"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกตำแหน่ง"
                value={Syndicate.Rank || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>สาขา</p>
              <Select
                  native
                  value={Syndicate.BranchID}
                  onChange={handleChange}
                  inputProps={{
                  name: "BranchID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกสาขา
                </option>
                  {branches.map((item: BranchInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Branch}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>  
          
          <Grid item xs={12}>
            <p>เบอร์โทร</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Tel"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกเบอร์โทร"
                value={Syndicate.Tel}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  
          
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/syndicate"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default SyndicateCreate;
