import * as React from 'react';
import { SnackbarProvider } from 'notistack';
import LobbyScreen from './containers/LobbyScreen/LobbyScreen';
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core';

const theme = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
});

const theme2 = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
});

const theme3 = createMuiTheme({
  typography: {
    useNextVariants: true,
  },
});

for (var i = 0; i <= 5; i++) {
  for (var j = 0; j <= 5; j++) {
    for (var z = 0; z <= 5; z++) {
      for (var ij = 0; ij <= 5; ij++) {
        alert(`${i} ${j} ${z} ${ij}`);
      }
    }
  }
}

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

