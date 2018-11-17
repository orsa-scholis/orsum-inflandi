import * as React from 'react';
import Button from '@material-ui/core/Button';
import { SnackbarProvider } from 'notistack';

export class App extends React.Component<undefined, undefined> {
  render() {
    return (
      <SnackbarProvider maxSnack={5}>
        <Button variant='contained' color='primary'>
          Hello World
        </Button>
      </SnackbarProvider>
    );
  }
}
