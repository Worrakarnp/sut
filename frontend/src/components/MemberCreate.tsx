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

import { MemberInterface } from "../models/IMember";
import { BranchInterface } from "../models/IBranch";
import { CategoryInterface } from "../models/ICategory";
import { SubDistrictInterface } from "../models/ISubDistrict";
import { DistrictInterface } from "../models/IDistrict";
import { ProvinceInterface } from "../models/IProvince";

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

function MemberCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [branches, setBranch] = useState<BranchInterface[]>([]);
  const [categories, setCategorys] = useState<CategoryInterface[]>([]);
  const [sub_districts, setSubDistrict] = useState<SubDistrictInterface[]>([]);
  const [districts, setDistrict] = useState<DistrictInterface[]>([]);
  const [provinces, setProvince] = useState<ProvinceInterface[]>([]);
  const [member, setMember] = useState<Partial<MemberInterface>>({});

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
    const name = event.target.name as keyof typeof member;
    setMember({
      ...member,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof member;
    const { value } = event.target;
    setMember({ ...member, [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
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

  const getCategory = async () => {
    fetch(`${apiUrl}/categories`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCategorys(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getSubDistrict = async () => {
    fetch(`${apiUrl}/sub_districts`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSubDistrict(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDistrict = async () => {
    fetch(`${apiUrl}/districts`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDistrict(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getProvince = async () => {
    fetch(`${apiUrl}/provinces`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setProvince(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getBranch();
    getCategory();
    getSubDistrict();
    getDistrict();
    getProvince();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      Name: member.Name ?? "",
      Num: member.Num,
      BranchID: convertType(member.BranchID),
      Email: member.Email,
      Tel: member.Tel,
      Address: member.Address,
      SubDistrictID: convertType(member.SubDistrictID),
      DistrictID: convertType(member.DistrictID),
      ProvinceID: convertType(member.ProvinceID),
      Date: selectedDate,
      CategoryID: convertType(member.CategoryID),
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

    fetch(`${apiUrl}/member`, requestOptionsPost)
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
              เพิ่มสมาชิกสมาคมฯ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={2} className={classes.root}>
          <Grid item xs={12}>
            <p>ชื่อสมาชิก</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกชื่อสมาชิก"
                value={member.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid> 

          <Grid item xs={12}>
            <p>เลขที่สมาชิก</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Num"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกเลขที่สมาชิก"
                value={member.Num}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>สาขา</p>
              <Select
                  native
                  value={member.BranchID}
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
            <p>อีเมล</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Email"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกอีเมล"
                value={member.Email}
                onChange={handleInputChange}
              />
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
                value={member.Tel}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>  

          <Grid item xs={12}>
            <p>ที่อยู่</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Address"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="โปรดกรอกที่อยู่"
                value={member.Address}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid> 

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>ตำบล</p>
              <Select
                  native
                  value={member.SubDistrictID}
                  onChange={handleChange}
                  inputProps={{
                  name: "SubDistrictID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกตำบล
                </option>
                  {sub_districts.map((item: SubDistrictInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.SubDistrict}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>อำเภอ</p>
              <Select
                  native
                  value={member.DistrictID}
                  onChange={handleChange}
                  inputProps={{
                  name: "DistrictID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกอำเภอ
                </option>
                  {districts.map((item: DistrictInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.District}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>จังหวัด</p>
              <Select
                  native
                  value={member.ProvinceID}
                  onChange={handleChange}
                  inputProps={{
                  name: "ProvinceID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกจังหวัด
                </option>
                  {provinces.map((item: ProvinceInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Province}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทสมาชิก</p>
              <Select
                  native
                  value={member.CategoryID}
                  onChange={handleChange}
                  inputProps={{
                  name: "CategoryID",
                  }}
                >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทสมาชิก
                </option>
                  {categories.map((item: CategoryInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Category}
                </option>
                  ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
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
              to="/member"
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

export default MemberCreate;
