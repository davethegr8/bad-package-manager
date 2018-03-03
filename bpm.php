#!/usr/bin/env php
<?php

define('VERSION', '0.1.2');

echo "Installing dependencies", PHP_EOL;

$dependencies = json_decode(file_get_contents('dependencies.json'), true);

foreach($dependencies['require'] as $dest => $src) {
	chdir(__DIR__);

	if(!is_dir($dest)) {
		mkdir($dest, 0777, true);
	}

	$is_empty = count(glob($dest.'/*')) === 0;
	$is_git = file_exists($dest.'/.git');

	$parts = explode('#', $src);
	$repo = $parts[0];
	if(isset($parts[1])) {
		$commitish = $parts[1];
	}
	else {
		$commitish = 'master';
	}

	chdir($dest);

	if($is_empty) {
		echo 'INFO Cloning "'.$src.'" into "./'.$dest.'"', PHP_EOL;
		run('git clone '.$repo.' .');
	}
	elseif(!$is_git) {
		echo 'WARNING folder "'.$dest.'" is not a git repository', PHP_EOL;
		throw new Exception("no git repository", 1);
	}
	else {
		$remotes = run('git remote -v');
		if(false === strpos($remotes, $repo)) {
			echo 'WARNING git remote src "'.$repo.'" does not match "'.substr($remotes, 0, strpos($remotes, "\n")).'" in folder "'.$dest.'"', PHP_EOL;
			throw new Exception("git remote mismatch", 1);
		}

		echo 'INFO Updating "'.$src.'" in "./'.$dest.'"', PHP_EOL;

		run('git fetch --all');
	}

	run('git checkout -q '.$commitish);
	run('git pull');
}

echo "Done.", PHP_EOL;

function run($cmd) {
	// echo $cmd, PHP_EOL;
	$output = shell_exec($cmd);
	return $output;
}