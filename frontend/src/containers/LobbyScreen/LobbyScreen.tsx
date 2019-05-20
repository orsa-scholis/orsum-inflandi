import * as React from 'react';
import * as PropTypes from 'prop-types';
import { AppBar, Grid, IconButton, Toolbar, Typography, withStyles } from '@material-ui/core';
import PlusIcon from '@material-ui/icons/Add';
import LobbyScreenStyles from './LobbyScreenStyles';
import { withSnackbar, withSnackbarProps } from 'notistack';
import GameList from '../../components/GameList/GameList';
import Game from '../../models/Game/Game';
import { History } from 'history';
import { Connection } from '../../connection/Connection';
import { Message } from "../../connection/Message";
import { Protocol } from "../../connection/protocol/Commands";

interface LobbyScreenProps extends withSnackbarProps {
  classes: any;
  history: History;
}

class LobbyScreen extends React.Component<LobbyScreenProps> {
  static propTypes = {
    classes: PropTypes.object,
  };

  componentDidMount(): void {
    const connection = new Connection('localhost', 4560, (err: Error) => console.log(err));
    connection.initiateHandshake();

    void connection.send(new Message(Protocol.ClientCommands.CONNECTION.CONNECT));
  }

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

export default withSnackbar(withStyles(LobbyScreenStyles)(LobbyScreen));
