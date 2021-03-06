# goObjStore

Simple Object Store in Golang including Peer to Peer and chunking algorithms to increase effiency.

[![Build Status](https://travis-ci.org/alruiz12/goObjStore.svg?branch=master)](https://travis-ci.org/alruiz12/goObjStore)[![codecov](https://codecov.io/gh/alruiz12/goObjStore/branch/master/graph/badge.svg)](https://codecov.io/gh/alruiz12/goObjStore)[![Code Health](https://landscape.io/github/alruiz12/goObjStore/master/landscape.svg?style=flat)](https://landscape.io/github/alruiz12/goObjStore/master)

## Architecture 

The project architecture is inspired by Openstack Swift as it has the elements Proxy and Storage Node, responsible for processing and storing data respectively. Unlike Swift, this project has a Tracker (inspired by the Bittorrent Protocol), instead of a ring, which links entities to their physical location.

![alt text](https://user-images.githubusercontent.com/22266492/30481223-c3cc18b6-9a1d-11e7-8d03-08fabbca81fb.PNG)

## Documentation

[Abstract](https://deim.urv.cat/~pfc/docs/pfc1548/d1504766046.pdf)

Complete [Documentation](https://deim.urv.cat/~pfc/docs/pfc1548/d1504766079.pdf) (Catalan Only)




