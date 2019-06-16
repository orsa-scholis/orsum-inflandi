import * as React from 'react';
import * as PropTypes from 'prop-types';
import { AppBar, Grid, IconButton, Toolbar, Typography, withStyles } from '@material-ui/core';
import PlusIcon from '@material-ui/icons/Add';
import LobbyScreenStyles from './LobbyScreenStyles';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import GameList from '../../components/GameList/GameList';
import Game from '../../models/Game/Game';
import { History } from 'history';
import { ServerConnection } from '../../connection/ServerConnection';

interface LobbyScreenProps extends WithSnackbarProps {
  classes: any;
  history: History;
}

class LobbyScreen extends React.Component<LobbyScreenProps> {
  static propTypes = {
    classes: PropTypes.object,
  };

  componentDidMount(): void {
    const connection = new ServerConnection('localhost', 4560, console.error);

    console.log('Sending connect');
    connection.sendConnect('Peter').then(answer => {
      console.log('received answer.');
      console.log(answer.payload!.getGamesList().map(game => game.getName()));
      console.dir(answer);
    }).catch(reason => {
      console.error(reason);
    });
  }

  // noinspection UnterminatedStatementJS
  gameSelected = (_game: Game) => {
    this.props.history.push('/game/1');
  }

  render() {
    const { classes, enqueueSnackbar } = this.props;

    const myGameList = [
      new Game('First game', 1),
      new Game('Second game', 2),
    ];

    return (
      <Grid container spacing={16}>
        <Grid item xs={12}>
          <AppBar position='static' color='default'>
            <Toolbar>
              <Typography variant='h6' color='textPrimary'>
                Lobby
              </Typography>
              <div className={classes.grow}/>
              <IconButton color='inherit' onClick={() => enqueueSnackbar('I am adding a game', { variant: 'success' })}>
                <PlusIcon/>
              </IconButton>
            </Toolbar>
          </AppBar>
        </Grid>
        <Grid item xs={12}>
          <GameList games={myGameList} onGameSelect={this.gameSelected} />
        </Grid>
      </Grid>
    );
  }
}

// @ts-ignore
export default withSnackbar(withStyles(LobbyScreenStyles)(LobbyScreen));
