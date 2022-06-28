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
import DateFnsUtils from '@date-io/date-fns';
import { KeyboardDateTimePicker, MuiPickersUtilsProvider } from "@material-ui/pickers";

import { GownInterface } from "../models/IGown";
import { SizeInterface } from "../models/ISize";
import { BlazerInterface } from "../models/IBlazer";
import { DegreeInterface } from "../models/IDegree";
import { BranchInterface } from "../models/IBranch";

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

function GownCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [sizes, setSize] = useState<SizeInterface[]>([]);
  const [blazers, setBlazer] = useState<BlazerInterface[]>([]);
  const [degrees, setDegree] = useState<DegreeInterface[]>([]);
  const [branches, setBranch] = useState<BranchInterface[]>([]);
  const [gown, setGown] = useState<Partial<GownInterface>>({});

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
    const name = event.target.name as keyof typeof gown;
    setGown({
      ...gown,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof gown;
    const { value } = event.target;
    setGown({ ...gown, [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getSize = async () => {
    fetch(`${apiUrl}/sizes`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSize(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getBlazer = async () => {
    fetch(`${apiUrl}/blazers`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setBlazer(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDegree = async () => {
    fetch(`${apiUrl}/degrees`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDegree(res.data);
        } else {
          console.log("else");
        }
      });
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
    getSize();
    getBlazer();
    getDegree();
    getBranch();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      Name: gown.Name ?? "",
      Student: gown.Student,
      Status: gown.Status,
      SizeID: convertType(gown.SizeID),
      BlazerID: convertType(gown.BlazerID),
      DegreeID:convertType(gown.DegreeID),
      BranchID: convertType(gown.BranchID),
      Tel: gown.Tel,
      Date: selectedDate,
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

    fetch(`${apiUrl}/gown`, requestOptionsPost)
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
              เพิ่มรายการ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={2} className={classes.root}>
          <Grid item xs={12}>
            <p>ชื่อ-นามสกุล</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกชื่อ-นามสกุล"
                value={gown.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid> 

          <Grid item xs={12}>
            <p>รหัสนักศึกษา</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Student"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกรหัสนักศึกษา"
                value={gown.Student}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  

          <Grid item xs={12}>
            <p>เช่า-ซื้อ-คืน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Status"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกเช่า-ซื้อ-คืน"
                value={gown.Status}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>ชุดครุยไซส์</p>
              <Select
                  native
                  value={gown.SizeID}
                  onChange={handleChange}
                  inputProps={{
                  name: "SizeID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกไซส์
                </option>
                  {sizes.map((item: SizeInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Size}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>เบลเซอร์</p>
              <Select
                  native
                  value={gown.BlazerID}
                  onChange={handleChange}
                  inputProps={{
                  name: "BlazerID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกเบลเซอร์
                </option>
                  {blazers.map((item: BlazerInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Blazer}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>ระดับปริญญา</p>
              <Select
                  native
                  value={gown.DegreeID}
                  onChange={handleChange}
                  inputProps={{
                  name: "DegreeID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกระดับปริญญา
                </option>
                  {degrees.map((item: DegreeInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Degree}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>สาขา</p>
              <Select
                  native
                  value={gown.BranchID}
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
                value={gown.Tel}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                    <p>วันที่บันทึก</p>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                    <KeyboardDateTimePicker
                        margin="normal"
                        id="date-picker-dialog"
                        label=""
                        format="dd/MM/yyyy"
                        value={selectedDate}
                        onChange={handleDateChange}
                        KeyboardButtonProps={{
                          'aria-label': 'change date',
                        }}
                    />
                </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/gown"
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

export default GownCreate;
