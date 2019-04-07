import * as React from 'react';
import { SnackbarProvider } from 'notistack';
import LobbyScreen from './containers/LobbyScreen/LobbyScreen';

export class App extends React.Component {
  render() {
    return (
      <SnackbarProvider maxSnack={5}>
        <LobbyScreen />
      </SnackbarProvider>
    );
  }
}
