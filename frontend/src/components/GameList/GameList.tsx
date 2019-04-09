import * as React from 'react';
import { List, ListItem, ListItemIcon, ListItemText } from '@material-ui/core';
import GameIcon from '@material-ui/icons/VideogameAsset';
import Game from '../../models/Game/Game';

interface GameListProps {
  games: Game[];
}

const GameListItem = ({ item }: { item: Game }) => {
  return (
    <ListItem button>
      <ListItemIcon>
        <GameIcon/>
      </ListItemIcon>
      <ListItemText primary={item.title}/>
    </ListItem>
  );
};

export default ({ games }: GameListProps) => {
  return (
    <div>
      <List>
        {games.map(game => <GameListItem item={game}/>)}
      </List>
    </div>
  );
};
