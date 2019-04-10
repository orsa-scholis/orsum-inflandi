import * as React from 'react';
import { SnackbarProvider } from 'notistack';
import LobbyScreen from './containers/LobbyScreen/LobbyScreen';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core';
import { Route, Router, Switch } from 'react-router-dom';
import { createHashHistory, History } from 'history';
import GameScreen from './containers/GameScreen/GameScreen';

const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
});

export class App extends React.Component<{}, { history: History }> {
  constructor(props: {}) {
    super(props);

    this.state = {
      history: createHashHistory()
    };
  }

  render() {
    return (
      <MuiThemeProvider theme={theme}>
        <SnackbarProvider maxSnack={5}>
          <Router history={this.state.history}>
            <Switch>
              <Route exact path='/game/:id' component={GameScreen}/>
              <Route component={LobbyScreen}/>
            </Switch>
          </Router>
        </SnackbarProvider>
      </MuiThemeProvider>
    );
  }
}
