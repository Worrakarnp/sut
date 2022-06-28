import React, { useEffect } from "react";
import clsx from "clsx";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import List from "@material-ui/core/List";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import Button from "@material-ui/core/Button";
import SearchIcon from "@material-ui/icons/Search";

//import FindInPageIcon from '@material-ui/icons/FindInPage';
import HomeIcon from "@material-ui/icons/Home";
//import VerticalAlignTopIcon from '@material-ui/icons/VerticalAlignTop';
import FaceIcon from '@material-ui/icons/Face';
import GroupIcon from '@material-ui/icons/Group';
import BusinessCenterIcon from '@material-ui/icons/BusinessCenter';
import SchoolIcon from '@material-ui/icons/School';
import Home from "./components/Home";
import SignIn from "./components/SignIn";
import Syndicate from "./components/Syndicate";
import SyndicateCreate from "./components/SyndicateCreate"; 
import Member from "./components/Member";
import MemberCreate from "./components/MemberCreate";
import Trader from "./components/Trader";
import TraderCreate from "./components/TraderCreate";
import Gown from "./components/Gown";
import GownCreate from "./components/GownCreate";
import UserFind from "./components/UserFind";
import MemberFind from "./components/MemberFind";
import TraderFind from "./components/TraderFind";
import GownFind from "./components/GownFind";

const drawerWidth = 240;
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },
    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
    a: {
      textDecoration: "none",
      color: "inherit",
    },
  })
);

export default function MiniDrawer() {
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const [token, setToken] = React.useState<String>("");
  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const menu = [
    { name: "หน้าหลัก", icon: <HomeIcon color="primary" />, path: "/" },
    { name: "ค้นหาคณะกรรมการสมาคมฯ", icon: <SearchIcon color="secondary" />, path: "/UserFind" },
    { name: "คณะกรรมการสมาคมฯ", icon: <FaceIcon />, path: "/syndicate" },
    { name: "ค้นหาสมาชิกสมาคมฯ", icon: <SearchIcon color="secondary" />, path: "/MemberFind" },
    { name: "สมาชิกสมาคมฯ", icon: <GroupIcon />, path: "/member" },
    { name: "ค้นหาผู้ประกอบการ", icon: <SearchIcon color="secondary" />, path: "/TraderFind" },
    { name: "ผู้ประกอบการ", icon: <BusinessCenterIcon />, path: "/trader" },
    { name: "ค้นหาชุดครุย", icon: <SearchIcon color="secondary" />, path: "/GownFind" },
    { name: "ระบบชุดครุย", icon: <SchoolIcon />, path: "/gown" },
  ];

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  return (
    <div className={classes.root}>
      <Router>
        <CssBaseline />
        {token && (
          <>
            <AppBar
              position="fixed"
              className={clsx(classes.appBar, {
                [classes.appBarShift]: open,
              })}
            >
              <Toolbar>
                <IconButton
                  color="inherit"
                  aria-label="open drawer"
                  onClick={handleDrawerOpen}
                  edge="start"
                  className={clsx(classes.menuButton, {
                    [classes.hide]: open,
                  })}
                >
                  <MenuIcon />
                </IconButton>
                <Typography variant="h6" className={classes.title}>
                ฐานข้อมูลภายในสมาคมเทคโนโลยีสุรนารี
                </Typography>
                <Button color="inherit" onClick={signout}>
                  Sign Out
                </Button>
              </Toolbar>
            </AppBar>
            <Drawer
              variant="permanent"
              className={clsx(classes.drawer, {
                [classes.drawerOpen]: open,
                [classes.drawerClose]: !open,
              })}
              classes={{
                paper: clsx({
                  [classes.drawerOpen]: open,
                  [classes.drawerClose]: !open,
                }),
              }}
            >
              <div className={classes.toolbar}>
                <IconButton onClick={handleDrawerClose}>
                  {theme.direction === "rtl" ? (
                    <ChevronRightIcon />
                  ) : (
                    <ChevronLeftIcon />
                  )}
                </IconButton>
              </div>
              <Divider />
              <List>
                {menu.map((item, index) => (
                  <Link to={item.path} key={item.name} className={classes.a}>
                    <ListItem button>
                      <ListItemIcon>{item.icon}</ListItemIcon>
                      <ListItemText primary={item.name} />
                    </ListItem>
                  </Link>
                ))}
              </List>
            </Drawer>
          </>
        )}

        <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/UserFind" component={UserFind} />
              <Route exact path="/syndicate" component={Syndicate} />
              <Route exact path="/syndicate/create" component={SyndicateCreate} />
              <Route exact path="/MemberFind" component={MemberFind} />
              <Route exact path="/member" component={Member} />
              <Route exact path="/member/create" component={MemberCreate} />
              <Route exact path="/TraderFind" component={TraderFind} />
              <Route exact path="/trader" component={Trader} />
              <Route exact path="/trader/create" component={TraderCreate} />
              <Route exact path="/GownFind" component={GownFind} />
              <Route exact path="/gown" component={Gown} />
              <Route exact path="/gown/create" component={GownCreate} />
	   </Switch>
          </div>
        </main>
      </Router>
    </div>
  );
}