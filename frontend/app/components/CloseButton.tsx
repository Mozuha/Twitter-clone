import { IoClose } from 'react-icons/io5';

import { Button } from '@components/material-tailwind';

export default function CloseButton() {
  return (
    <Button variant="text" color="gray" className="w-9 h-9 pt-1.5 pl-1.5" aria-label="close-button">
      <IoClose color="white" className="w-6 h-6" />
    </Button>
  );
}
