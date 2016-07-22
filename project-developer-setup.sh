#!/bin/bash

OS=`uname`

if [ "$OS"="Darwin" ] ; then 
		#Ensure HomeBrew is installed.  This is the package manager of choice for Mac OS
	command -v brew >/dev/null 2>&1 || { 
		ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	}

	#Is Go already installed?
	command -v go >/dev/null 2>&1 || { echo >&2 "Installing Go from Homebrew"; brew install go --cross-compile-common; exit 1; }

	#Sublime is the editor of choice for Golang programmers.
	command -v subl >/dev/null 2>&1 || { 
		brew install caskroom/cask/brew-cask
		brew tap caskroom/versions
		brew cask install sublime-text3

		echo "you must also link the 'subl' command to your path: sudo ln -s \"/Applications/Sublime Text 3.app/Contents/SharedSupport/bin/subl\" /usr/local/bin/subl"; read -p "Press [Enter] Once the Above is complete..."; echo "Install the golang plugin for Sublime Text 2 as documented here: https://github.com/DisposaBoy/GoSublime"; read -p "Press [Enter] Once the Above Plugin Is Installed"; 
	}
else
			
	#Ensure API is installed.  This is the package manager we support
	command -v apt-get >/dev/null 2>&1 || { echo >&2 "Aptitude is required for these setup scripts but it's not installed.  Aborting."; exit 1; }


	#Is Go already installed?
	command -v go >/dev/null 2>&1 || { echo >&2 "Installing Go from Aptitude";sudo apt-get install golang; exit 1; }


	#install gvm
	command -v gvm >/dev/null 2>&1 || { echo >&2 "Installing GVM for better management of Go Versions" 
		sudo apt-get install bison
		curl https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
		source /home/pwilson/.gvm/scripts/gvm

		#create the compile directory
		gvm install go1.4
		gvm use go1.4  #needed as go1.5 needs to be compiled with go1.4.
		gvm install go1.5
		gvm use go1.5
	}

	#Sublime is the editor of choice for Golang programmers.
	command -v subl >/dev/null 2>&1 || { 
		echo >&2 "Installing Sublime Text 3."
		sudo add-apt-repository ppa:webupd8team/sublime-text-3
		sudo apt-get update
		sudo apt-get install sublime-text-installer

		echo "Install the golang plugin for Sublime Text 3 as documented here: https://github.com/DisposaBoy/GoSublime"
		read -p "Press [Enter] Once the Above Plugin Is Installed"; 
	}
fi