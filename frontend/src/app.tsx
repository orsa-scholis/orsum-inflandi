import * as React from 'react';
import { SnackbarProvider } from 'notistack';
import LobbyScreen from './containers/LobbyScreen/LobbyScreen';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core';

const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
});

export class App extends React.Component {
  render() {
    return (
      <MuiThemeProvider theme={theme}>
        <SnackbarProvider maxSnack={5}>
          <LobbyScreen />
        </SnackbarProvider>
      </MuiThemeProvider>
    );
  }
}
