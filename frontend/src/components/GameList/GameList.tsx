import * as React from 'react';
import { List, ListItem, ListItemIcon, ListItemText } from '@material-ui/core';
import GameIcon from '@material-ui/icons/VideogameAsset';
import Game from '../../models/Game/Game';

interface GameListProps {
  games: Game[];
  onGameSelect: (game: Game) => void;
}

const GameListItem = ({ item, onClick }: { item: Game, onClick?: () => void }) => {
  return (
    <ListItem onClick={onClick} button>
      <ListItemIcon>
        <GameIcon/>
      </ListItemIcon>
      <ListItemText primary={item.title}/>
    </ListItem>
  );
};

export default ({ games, onGameSelect }: GameListProps) => {
  return (
    <List>
      {games.map(game => (
        <GameListItem key={game.id} item={game} onClick={() => onGameSelect(game)}/>
      ))}
    </List>
  );
};
