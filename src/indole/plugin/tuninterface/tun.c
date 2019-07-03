#ifdef __MINGW64__
#include <stdio.h>
#include <stdlib.h>
int32_t setup_tap_device(char *ifname) {
  perror("tun/tap is not support on windows");
  return -1;
}

int32_t setup_tun_device(char *ifname) {
  perror("tun/tap is not support on windows");
  return -1;
}

int32_t up_device(char *name) {
  perror("tun/tap is not support on windows");
  return -1;
}

int32_t set_ip(char *name, char *ip_addr, char *netmask) {
  perror("tun/tap is not support on windows");
  return -1;
}
#endif

#ifdef __linux__
#include <arpa/inet.h>
#include <fcntl.h>
#include <linux/if.h>
#include <linux/if_tun.h>
#include <net/route.h>
#include <netinet/in.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
int32_t setup_dev(char *ifname, short flags) {
  struct ifreq ifr;
  int err,fd;
  if ((fd = open("/dev/net/tun", O_RDWR)) < 0) {
    return fd;
  }
  memset(&ifr, 0, sizeof(ifr));
  ifr.ifr_flags = flags;
  strncpy(ifr.ifr_name, ifname, IFNAMSIZ);
  if ((err = ioctl(fd, TUNSETIFF, &ifr)) < 0) {
    close(fd);
    return err;
  }
  strncpy(ifname, ifr.ifr_name, IFNAMSIZ);
  return fd;
}

int32_t setup_tap_device(char *ifname) {
  return setup_dev(ifname, IFF_TAP | IFF_NO_PI);
}

int32_t setup_tun_device(char *ifname) {
  return setup_dev(ifname, IFF_TUN | IFF_NO_PI);
}

int32_t up_device(char *name) {
  struct ifreq ifr;
  int sockfd;

  if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
    perror("Create socket fails!\n");
    return -1;
  }

  strncpy(ifr.ifr_name, name, IFNAMSIZ);
  if (ioctl(sockfd, SIOCGIFFLAGS, &ifr) < 0) {
    perror("ioctl SIOCGIFFLAGS fails!\n");
    close(sockfd);
    return -1;
  }

  ifr.ifr_flags |= IFF_UP;
  if (ioctl(sockfd, SIOCSIFFLAGS, &ifr) < 0) {
    perror("ioctl SIOCSIFFLAGS fails!\n");
    close(sockfd);
    return -1;
  }

  close(sockfd);

  return 1;
}

int32_t set_ip(char *name, char *ip_addr, char *netmask) {
  up_device(name);
  perror(ip_addr);
  perror(netmask);
  int sockfd;
  if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
    perror("Create socket fails!\n");
    return -1;
  }
  struct ifreq ifr;
  struct sockaddr_in sin;
  strncpy(ifr.ifr_name, name, IFNAMSIZ);
  if (ioctl(sockfd, SIOCSIFFLAGS, &ifr) < 0) {
    return -4;
  }
  sin.sin_family = AF_INET;
  inet_aton(ip_addr, &(sin.sin_addr));
  memcpy(&ifr.ifr_addr, &sin, sizeof(struct sockaddr));
  if (ioctl(sockfd, SIOCSIFADDR, &ifr) < 0) {
    return -2;
  }
  inet_aton(netmask, &(sin.sin_addr));
  memcpy(&ifr.ifr_netmask, &sin, sizeof(struct sockaddr));
  if (ioctl(sockfd, SIOCSIFNETMASK, &ifr) < 0) {
    return -3;
  }
  return 1;
}

int32_t set_mtu(char *name,int32_t mtu) {
  printf("%d",mtu);
  int sockfd;
  if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
    perror("Create socket fails!\n");
    return -1;
  }
  struct ifreq ifr;
  struct sockaddr_in sin;
  strncpy(ifr.ifr_name, name, IFNAMSIZ);
  if (ioctl(sockfd, SIOCSIFFLAGS, &ifr) < 0) {
    return -4;
  }
  ifr.ifr_mtu = mtu; 
  if (ioctl(sockfd, SIOCSIFMTU, &ifr) < 0) {
    return -2;
  }
  return 1;
}
#endif

// int main() {
//     int err;
//     char dev[IFNAMSIZ];
//     memset(&dev,0,sizeof(dev));
//     strncpy(dev,"tun1",4);
//     if (setup_tun_device(dev) < 0){
//         printf("setup failed\n");
//         return -1;
//     }
//     if (up_device(dev) < 0){
//         printf("up failed\n");
//         return -1;
//     }
//     if (set_ip(dev,"192.168.1.2","255.255.255.0") < 0){
//         printf("set_ip failed\n");
//         return -1;
//     }
//     if (set_mtu(dev,1420) < 0){
//         printf("set_mtu failed\n");
//         return -1;
//     }
//     sleep(10);
//     return 0;
// }
