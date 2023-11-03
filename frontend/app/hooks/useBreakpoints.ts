import useMediaQuery from './useMediaQuery';

// https://dev.to/justincy/4-patterns-for-responsive-props-in-react-39ak
export default function useBreakpoints() {
  const breakpoints = {
    isSm: useMediaQuery('(max-width: 639.98px)'),
    isMd: useMediaQuery('(min-width: 640px) and (max-width: 767.98px)'),
    isLg: useMediaQuery('(min-width: 768px) and (max-width: 1023.98px)'),
    isXl: useMediaQuery('(min-width: 1024px)'),
    active: 'lg',
  };

  if (breakpoints.isSm) breakpoints.active = 'sm';
  if (breakpoints.isMd) breakpoints.active = 'md';
  if (breakpoints.isLg) breakpoints.active = 'lg';
  if (breakpoints.isXl) breakpoints.active = 'xl';

  return breakpoints;
}
